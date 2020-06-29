package service

import (
	coma "github.com/SAIKAII/go-conn-manager"
	"github.com/SAIKAII/skHappy-IM/infra/base"
	"github.com/SAIKAII/skHappy-IM/internal/logic/cache"
	"github.com/SAIKAII/skHappy-IM/internal/logic/dao"
	pb "github.com/SAIKAII/skHappy-IM/protocols"
	"github.com/SAIKAII/skHappy-IM/services"
	"github.com/golang/protobuf/proto"
	"time"
)

type messageService struct {
}

func (ms *messageService) SendToOne(req *pb.DeliverMessageReq) error {
	conn := base.ConnectionManager().GetConn(req.Item.ReceiverName)
	if conn == nil {
		// 对方不在线，保存消息到数据库后直接返回
		return nil
	}

	output := &pb.MessageOutput{
		Item:  req.Item,
		SeqId: req.SeqId,
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

func (ms *messageService) SaveMessage(req *pb.SendMessageReq) (uint64, error) {
	typ, content := services.PBToContent(req.Item.MsgBody)
	tm := time.Unix(req.Item.SendTime, 0)
	key := cache.UserKey(req.Item.ReceiverName)
	seqId, err := cache.Incr(key)
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
		cache.UpdateUserSeq(req.Item.ReceiverName, seqId)
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
	// TODO 两个数据库操作做成事务，失败的话要DECR seqId
	err = messageDao.InsertOne(msg)
	if err != nil {

	}
	err = msgRecvDao.UpdateLastSeqId(req.Item.ReceiverName, seqId)
	if err != nil {

	}

	return seqId, nil
}
