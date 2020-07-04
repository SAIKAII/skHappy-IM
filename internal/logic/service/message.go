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
		return err
	}

	return nil
}

func (ms *messageService) SendToOne(ctx context.Context, req *pb.SendMessageReq) error {
	// 先判断两人是否好友关系
	isFriend, err := services.IRelationshipService.IsFriend(req.Item.SenderName, req.Item.ReceiverName)
	if err != nil {
		return err
	}

	if !isFriend {
		return errors.New("两人不是好友关系")
	}

	err = ms.sendToPeer(ctx, req)
	if err != nil {
		return err
	}

	return nil
}

func (ms *messageService) SendToGroup(ctx context.Context, req *pb.SendMessageReq) error {
	// 先判断该用户是否群员
	isMember, err := services.IGroupService.IsMember(req.Item.GroupId, req.Item.SenderName)
	if err != nil {
		return err
	}

	if !isMember {
		return errors.New("该用户不在群组内")
	}

	users, err := services.IGroupService.ListGroupMember(req.Item.GroupId)
	if err != nil {
		return err
	}

	for _, u := range users {
		if req.Item.SenderName == u.Username {
			continue
		}
		req.Item.ReceiverName = u.Username
		err := ms.sendToPeer(ctx, req)
		if _, ok := err.(*mysql.MySQLError); ok {
			// TODO 数据库错误，也就是没有正确保存，需要处理
		}
	}

	return nil
}

func (ms *messageService) sendToPeer(ctx context.Context, req *pb.SendMessageReq) error {
	seqId, err := services.IMessageService.SaveMessage(ctx, req)
	req.Item.SeqId = seqId
	if err != nil {
		return err
	}

	rpcCli := base.NewRPCCli()
	rpcConn, err := rpcCli.Dialer(base.USER_ADDR, req.Item.ReceiverName)
	if err != nil {
		return err
	}
	_, err = pb.NewConnServiceClient(rpcConn).DeliverMessage(ctx, &pb.DeliverMessageReq{
		Item: req.Item,
	})

	if err != nil && err != redis.ErrNil {
		return err
	}

	return nil
}

func (ms *messageService) SendToUser(ctx context.Context, req *pb.DeliverMessageReq) error {
	conn := base.ConnectionManager().GetConn(req.Item.ReceiverName)
	if conn == nil {
		// 对方不在线，保存消息到数据库后直接返回
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
		return err
	}

	return nil
}

func (ms *messageService) SaveMessage(ctx context.Context, req *pb.SendMessageReq) (uint64, error) {
	typ, content := services.PBToContent(req.Item.MsgBody)
	tm := time.Unix(req.Item.SendTime, 0)
	key := cache.SeqCache.Key(req.Item.ReceiverName)
	seqId, err := cache.SeqCache.Incr(key)
	if err != nil {
		return 0, err
	}

	db := base.Database()
	msgRecvDao := dao.MsgRecvDao{DB: db}
	// 如果seqId为1,可能数据库有更加新的数据，所以尝试从数据库取该用户的seqId，然后更新缓存中的seqId
	if seqId == 1 {
		msgRecv, err := msgRecvDao.GetOne(req.Item.ReceiverName)
		if err != nil {
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

	err = db.Transaction(func(tx *gorm.DB) error {
		messageDao := dao.MessageDao{DB: tx}
		err = messageDao.InsertOne(msg)
		if err != nil {
			return err
		}
		recvDao := dao.MsgRecvDao{DB: tx}
		err = recvDao.UpdateLastSeqId(req.Item.ReceiverName, seqId)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		// 事务失败，seqId恢复原来的值
		seqId, e := cache.SeqCache.Decr(key)
		if e != nil {
			return 0, err
		}
		msgRecvDao.UpdateLastSeqId(req.Item.ReceiverName, seqId)
		return 0, err
	}

	return seqId, nil
}
