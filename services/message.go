package services

import (
	"context"
	"github.com/SAIKAII/skHappy-IM/pkg/util"
	pb "github.com/SAIKAII/skHappy-IM/protocols"
	jsoniter "github.com/json-iterator/go"
)

type MessageType int8

var IMessageService MessageService

type MessageService interface {
	Send(context.Context, *pb.SendMessageReq) error
	SendToOne(context.Context, *pb.SendMessageReq) error
	SendToGroup(context.Context, *pb.SendMessageReq) error
	SendToUser(context.Context, *pb.DeliverMessageReq) error
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
