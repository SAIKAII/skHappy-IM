package service

import (
	"context"
	"errors"
	coma "github.com/SAIKAII/go-conn-manager"
	"github.com/SAIKAII/skHappy-IM/infra/base"
	"github.com/SAIKAII/skHappy-IM/internal/logic/cache"
	"github.com/SAIKAII/skHappy-IM/internal/logic/dao"
	pb "github.com/SAIKAII/skHappy-IM/protocols"
	"github.com/SAIKAII/skHappy-IM/services"
	"github.com/go-sql-driver/mysql"
	"github.com/golang/protobuf/proto"
	"github.com/gomodule/redigo/redis"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
	"time"
)

type messageService struct {
}

func (ms *messageService) Send(ctx context.Context, req *pb.SendMessageReq) error {
	var err error
	switch req.Item.ReceiverType {
	case pb.ReceiverType_RT_USER:
		err = ms.SendToOne(ctx, req)
	case pb.ReceiverType_RT_GROUP:
		err = ms.SendToGroup(ctx, req)
	}

	if err != nil {
		base.Logger.Errorln(err)
		return err
	}

	return nil
}

func (ms *messageService) SendToOne(ctx context.Context, req *pb.SendMessageReq) error {
	// 先判断两人是否好友关系
	isFriend, err := services.IRelationshipService.IsFriend(req.Item.SenderName, req.Item.ReceiverName)
	if err != nil {
		base.Logger.Errorln(err)
		return err
	}

	if !isFriend {
		err := errors.New("两人不是好友关系")
		base.Logger.Errorln(err)
		return err
	}

	db := base.Database()
	seqId, err := ms.saveMessage(req, db)
	if _, ok := err.(*mysql.MySQLError); ok {
		// TODO 数据库错误，也就是没有正确保存，需要处理
	}
	req.Item.SeqId = seqId
	err = ms.sendToPeer(ctx, req)
	if err != nil {
		base.Logger.Errorln(err)
		return err
	}

	return nil
}

func (ms *messageService) SendToGroup(ctx context.Context, req *pb.SendMessageReq) error {
	// 先判断该用户是否群员
	isMember, err := services.IGroupService.IsMember(req.Item.GroupId, req.Item.SenderName)
	if err != nil {
		base.Logger.Errorln(err)
		return err
	}

	if !isMember {
		err := errors.New("该用户不在群组内")
		base.Logger.Errorln(err)
		return err
	}

	users, err := services.IGroupService.ListGroupMember(req.Item.GroupId)
	if err != nil {
		base.Logger.Errorln(err)
		return err
	}

	// 保存到数据库
	seqIds := make([]uint64, len(users))
	db := base.Database()
	logrus.Info(time.Now(), "开始保存消息")
	err = db.Transaction(func(tx *gorm.DB) error {
		for i, u := range users {
			if req.Item.SenderName == u.Username {
				continue
			}
			req.Item.ReceiverName = u.Username
			seqId, err := ms.saveMessage(req, tx)
			if err != nil {
				if _, ok := err.(*mysql.MySQLError); ok {
					// TODO 数据库错误，也就是没有正确保存，需要处理
					base.Logger.Errorln(err)
				}
			}

			seqIds[i] = seqId
		}
		return nil
	})
	if err != nil {
		base.Logger.Errorln(err)
		// 消息保存失败，不发送消息
		return err
	}
	logrus.Info(time.Now(), "保存完毕")

	// 发送给在线用户
	for i := range users {
		if req.Item.SenderName == users[i].Username {
			continue
		}
		req.Item.SeqId = seqIds[i]
		req.Item.ReceiverName = users[i].Username
		err = ms.sendToPeer(ctx, req)
		if err != nil {
			// 发送给某个用户失败，但数据库有保存该消息，暂不处理
			base.Logger.Errorln(err)
		}
	}

	return nil
}

func (ms *messageService) sendToPeer(ctx context.Context, req *pb.SendMessageReq) error {
	rpcCli := base.NewRPCCli()
	rpcConn, err := rpcCli.Dialer(base.USER_ADDR, req.Item.ReceiverName)
	if err != nil {
		base.Logger.Errorln(err)
		return err
	}
	_, err = pb.NewConnServiceClient(rpcConn).DeliverMessage(ctx, &pb.DeliverMessageReq{
		Item: req.Item,
	})

	if err != nil && err != redis.ErrNil {
		base.Logger.Errorln(err)
		return err
	}

	return nil
}

func (ms *messageService) SendToUser(ctx context.Context, req *pb.DeliverMessageReq) error {
	conn := base.ConnectionManager().GetConn(req.Item.ReceiverName)
	if conn == nil {
		// 操作到达这一步骤，对方却不在线，说明缓存中有过期数据
		rdConn := base.RedisConn()
		defer rdConn.Close()
		rdConn.Do("HDEL", base.USER_ADDR, req.Item.ReceiverName)
		return nil
	}

	output := &pb.MessageOutput{
		Item: req.Item,
	}
	b, _ := proto.Marshal(output)
	dmReq := &pb.ConnOutput{
		PackageType: pb.PackageType_PT_MESSAGE,
		ErrCode:     0,
		ErrMsg:      "",
		Data:        b,
	}
	o, _ := proto.Marshal(dmReq)
	err := coma.PacketToPeer(conn, o)
	if err != nil {
		base.Logger.Errorln(err)
		return err
	}

	return nil
}

func (ms *messageService) saveMessage(req *pb.SendMessageReq, db *gorm.DB) (uint64, error) {
	typ, content := services.PBToContent(req.Item.MsgBody)
	tm := time.Unix(req.Item.SendTime, 0)
	key := cache.SeqCache.Key(req.Item.ReceiverName)
	seqId, err := cache.SeqCache.Incr(key)
	if err != nil {
		base.Logger.Errorln(err)
		return 0, err
	}

	msgRecvDao := dao.MsgRecvDao{DB: db}
	// 如果seqId为1,可能数据库有更加新的数据，所以尝试从数据库取该用户的seqId，然后更新缓存中的seqId
	if seqId == 1 {
		msgRecv, err := msgRecvDao.GetOne(req.Item.ReceiverName)
		if err != nil {
			base.Logger.Errorln(err)
			return 0, err
		}

		seqId = msgRecv.LastSeqId
		seqId++
		cache.SeqCache.UpdateUserSeq(req.Item.ReceiverName, seqId)
	}

	msg := &dao.Message{
		SenderType:   int8(req.Item.SenderType),
		Sender:       req.Item.SenderName,
		ReceiverType: int8(req.Item.ReceiverType),
		Receiver:     req.Item.ReceiverName,
		Type:         typ,
		Content:      content,
		SeqId:        seqId,
		SendTime:     &tm,
	}

	messageDao := dao.MessageDao{DB: db}
	err = messageDao.InsertOne(msg)
	if err != nil {
		base.Logger.Errorln(err)
		return 0, err
	}
	recvDao := dao.MsgRecvDao{DB: db}
	err = recvDao.UpdateLastSeqId(req.Item.ReceiverName, seqId)
	if err != nil {
		// 事务失败，seqId恢复原来的值
		seqId, e := cache.SeqCache.Decr(key)
		if e != nil {
			base.Logger.Errorln(err)
			return 0, err
		}
		msgRecvDao.UpdateLastSeqId(req.Item.ReceiverName, seqId)
		return 0, err
	}

	return seqId, nil
}
