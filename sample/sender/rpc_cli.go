package main

import (
	"context"
	"fmt"
	pb "github.com/SAIKAII/skHappy-IM/protocols"
	"time"
)

func register(username, nickname string) {
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
	if err != nil {
		panic(err)
	}
}

func addFriend(self, friend string) {
	req := &pb.AddFriendReq{
		UserId:   self,
		FriendId: friend,
	}
	_, err := cc.AddFriend(context.TODO(), req)
	if err != nil {
		panic(err)
	}
}

func getFriend(username string) {
	req := &pb.GetUserReq{
		Username: username,
	}
	resp, err := cc.GetFriend(context.TODO(), req)
	if err != nil {
		panic(err)
	}
	fmt.Println("[GetFriend]=>", resp.User)
}

func listFriends(username string) {
	req := &pb.ListUsersReq{
		Username: username,
	}
	resp, err := cc.ListFriends(context.TODO(), req)
	if err != nil {
		panic(err)
	}
	fmt.Println("[ListFriends]=>", resp.Users)
}

func deleteFriend(self, friend string) {
	req := &pb.DelFriendReq{
		Username:   self,
		FriendName: friend,
	}
	_, err := cc.DelFriend(context.TODO(), req)
	if err != nil {
		panic(err)
	}
}

func updateProfile(username, nickname string) {
	user := &pb.User{
		Username: username,
		Nickname: nickname,
		Password: "123456",
		Sex:      0,
		Birthday: time.Date(1995, time.January, 10, 0, 0, 0, 0, time.UTC).Unix(),
		PhoneNum: "12473762905",
	}
	req := &pb.UpdateProfileReq{
		User: user,
	}
	_, err := cc.UpdateProfile(context.TODO(), req)
	if err != nil {
		panic(err)
	}
}

func changePassword(username, now string) {
	req := &pb.ChangePasswordReq{
		Username:    username,
		PrePassword: "123456",
		Password:    now,
	}
	_, err := cc.ChangePassword(context.TODO(), req)
	if err != nil {
		panic(err)
	}
}

func sendMessage(sender, receiver string) {
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
		ReceiverName: receiver,
		ReceiverType: pb.ReceiverType_RT_USER,
		MsgBody:      msg,
		SendTime:     time.Now().Unix(),
	}
	req := &pb.SendMessageReq{
		Item: item,
	}
	_, err := cc.SendMessage(context.TODO(), req)
	if err != nil {
		panic(err)
	}
}
