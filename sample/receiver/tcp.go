package main

import (
	"context"
	"fmt"
	pb "github.com/SAIKAII/skHappy-IM/protocols"
	"github.com/golang/protobuf/proto"
	"time"
)

func getMessage(data []byte) error {
	var content pb.MessageOutput
	err := proto.Unmarshal(data, &content)
	if err != nil {
		return err
	}

	fmt.Println("[From]", content.Item.SenderName, "[To]", content.Item.ReceiverName,
		"[Type]", content.Item.MsgBody.Type, "[content]", content.Item.MsgBody.Content.GetText().Text)

	retMessage(content)
	return nil
}

func retMessage(resp pb.MessageOutput) error {
	t := &pb.Text{
		Text: resp.Item.MsgBody.Content.GetText().Text + "-",
	}
	ct := &pb.MessageContent_Text{
		Text: t,
	}
	mc := &pb.MessageContent{
		Content: ct,
	}
	msg := &pb.MessageBody{
		Type:    pb.MessageType_MT_TEXT,
		Content: mc,
	}
	item := &pb.MessageItem{
		SenderName:   resp.Item.ReceiverName,
		SenderType:   pb.SenderType_ST_USER,
		ReceiverName: resp.Item.SenderName,
		ReceiverType: pb.ReceiverType_RT_USER,
		MsgBody:      msg,
		SendTime:     time.Now().Unix(),
	}
	req := &pb.SendMessageReq{
		Item: item,
	}

	_, err := cc.SendMessage(context.TODO(), req)
	if err != nil {
		return err
	}

	return nil
}
