package main

import (
	"context"
	pb "github.com/SAIKAII/skHappy-IM/protocols"
	"github.com/SAIKAII/skHappy-IM/sample/util"
	"time"
)

func register(username, nickname string) error {
	user := &pb.User{
		Username: username,
		Nickname: nickname,
		Password: "123456",
		Sex:      0,
		Birthday: time.Date(1995, time.January, 10, 0, 0, 0, 0, time.UTC).Unix(),
		PhoneNum: "12473762905",
	}
	in := &pb.RegisterReq{
		User: user,
	}
	_, err := cc.Register(context.TODO(), in)
	return err
}

func createGroup(username string) (uint64, error) {
	groupName := util.RandString(30)
	group := &pb.Group{
		GroupName:    groupName,
		CreateUser:   username,
		Owner:        username,
		Announcement: "",
	}
	req := &pb.CreateGroupReq{
		Group: group,
	}
	resp, err := cc.CreateGroup(context.TODO(), req)
	if err != nil {
		return 0, err
	}

	return resp.GroupId, nil
}

func sendGroupMessage(sender string, groupId uint64) error {
	t := &pb.Text{
		Text: "Hello world",
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
		SenderName:   sender,
		SenderType:   pb.SenderType_ST_USER,
		GroupId:      groupId,
		ReceiverType: pb.ReceiverType_RT_GROUP,
		MsgBody:      msg,
		SendTime:     time.Now().Unix(),
	}
	req := &pb.SendMessageReq{
		Item: item,
	}
	_, err := cc.SendMessage(context.TODO(), req)
	return err
}
