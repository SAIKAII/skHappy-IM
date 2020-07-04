package main

import (
	"context"
	"fmt"
	pb "github.com/SAIKAII/skHappy-IM/protocols"
	"github.com/SAIKAII/skHappy-IM/sample/long_link"
	codec "github.com/SAIKAII/skHappy-IM/sample/util"
	"github.com/golang/protobuf/proto"
	"net"
	"time"
)

func tcpConnTest(selfName string, c chan interface{}) {
	conn, err := net.Dial("tcp", "127.0.0.1:8090")
	defer conn.Close()
	if err != nil {
		panic(err)
	}

	cdc := codec.NewCodec(conn)
	defer cdc.CloseConnection()
	go long_link.ReadResp(cdc, getMessage, c)

	// 登录
	req := &pb.SignInReq{
		Username: selfName,
		Password: "123456",
	}
	err = long_link.Login(cdc, req)
	if err != nil {
		panic(err)
	}

	long_link.HeartBeat(cdc)
}

func getMessage(data []byte) error {
	var content pb.MessageOutput
	err := proto.Unmarshal(data, &content)
	if err != nil {
		return err
	}

	fmt.Println("[From]", content.Item.SenderName, "[To]", content.Item.ReceiverName,
		"[Type]", content.Item.MsgBody.Type, "[content]", content.Item.MsgBody.Content.GetText().Text)

	//retMessage(content)
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
		GroupId:      resp.Item.GroupId,
		ReceiverType: pb.ReceiverType_RT_GROUP,
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
