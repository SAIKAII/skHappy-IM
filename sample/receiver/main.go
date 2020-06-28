package main

import (
	"context"
	"fmt"
	pb "github.com/SAIKAII/skHappy-IM/protocols"
	"github.com/SAIKAII/skHappy-IM/sample/long_link"
	codec "github.com/SAIKAII/skHappy-IM/sample/util"
	"github.com/golang/protobuf/proto"
	"google.golang.org/grpc"
	"net"
	"os"
	"os/signal"
	"time"
)

var cc pb.CliInterfaceServiceClient

func main() {
	cli, err := grpc.Dial("127.0.0.1:8088", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	cc = pb.NewCliInterfaceServiceClient(cli)

	conn, err := net.Dial("tcp", "127.0.0.1:8090")
	if err != nil {
		panic(err)
	}

	cdc := codec.NewCodec(conn)
	go long_link.ReadResp(cdc, getMessage)

	// 登录
	req := &pb.SignInReq{
		Username: "qffqwrtb231",
		Password: "890567",
	}
	err = long_link.Login(cdc, req)
	if err != nil {
		panic(err)
	}

	go long_link.HeartBeat(cdc)

	// 拉取离线消息
	sync := &pb.SyncReq{
		Username:  "qffqwrtb231",
		LastSeqId: 70,
	}
	err = long_link.Sync(cdc, sync)
	if err != nil {
		panic(err)
	}
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, os.Kill)
	select {
	case <-c:
		cdc.CloseConnection()
		fmt.Println("Stop")
	}
}

func getMessage(data []byte) error {
	var content pb.MessageOutput
	err := proto.Unmarshal(data, &content)
	if err != nil {
		return err
	}

	fmt.Println("[From]", content.Item.SenderName, "[To]", content.Item.ReceiverName,
		"[Type]", content.Item.MsgBody.Type, "[content]", content.Item.MsgBody.Content.GetText().Text)

	retMessage(content)
	//time.Sleep(1 * time.Second)
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
