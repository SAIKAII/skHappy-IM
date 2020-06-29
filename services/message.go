package services

import (
	"github.com/SAIKAII/skHappy-IM/pkg/util"
	pb "github.com/SAIKAII/skHappy-IM/protocols"
	jsoniter "github.com/json-iterator/go"
)

type MessageType int8

const (
	MESSAGE_TYPE_SINGLE     MessageType = iota // 单对单消息
	MESSAGE_TYPE_GROUP                         // 群聊消息
	MESSAGE_TYPE_HEART_BEAT                    // 心跳

)

const REDIS_MESSAGE_ID string = "msgIdMap" // redis中保存seqId的map名

const (
	MESSAGE_NOT_DELIVERED = iota // 消息未送达
	MESSAGE_DELIVERED            // 消息送达
)

var IMessageService MessageService

type MessageService interface {
	SendToOne(*pb.DeliverMessageReq) error
	//SendToGroup(MessageTransferDTO) error
	SaveMessage(*pb.SendMessageReq) (uint64, error)
}

type MessageTransferDTO struct {
	MsgFrom string
	MsgTo   string
	MsgType MessageType
	Content []byte
}

func PBToContent(msgBody *pb.MessageBody) (int8, string) {
	var content interface{}
	switch msgBody.Type {
	case pb.MessageType_MT_IMAGE:
		content = msgBody.Content.GetImage()
	case pb.MessageType_MT_TEXT:
		content = msgBody.Content.GetText()
	}

	b, _ := jsoniter.Marshal(content)
	return int8(msgBody.Type), util.Bytes2Str(b)
}
