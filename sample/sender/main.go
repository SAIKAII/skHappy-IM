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
	"time"
)

var cc pb.CliInterfaceServiceClient

func main() {
	cli, err := grpc.Dial("127.0.0.1:8088", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	cc = pb.NewCliInterfaceServiceClient(cli)
	// 注册帐号
	//register()
	// 添加好友
	//addFriend()
	// 获取好友信息
	getFriend()
	// 列出所有好友
	listFriends()
	// 删除好友关系
	//deleteFriend()
	// 更新个人信息
	//updateProfile()
	// 更换密码
	//changePassword()
	// 发送消息
	sendMessage()

	// tcp测试
	tcpConnTest()
}

func register() {
	user := &pb.User{
		Username: "qffqwrtb231",
		Nickname: "SAIKAII",
		Password: "123456",
		Sex:      0,
		Birthday: time.Date(1995, time.January, 10, 0, 0, 0, 0, time.UTC).Unix(),
		PhoneNum: "12405762905",
	}
	in := &pb.RegisterReq{
		User: user,
	}
	_, err := cc.Register(context.TODO(), in)
	if err != nil {
		panic(err)
	}
}

func addFriend() {
	friend := &pb.AddFriendReq{
		UserId:   "qffqwrtb231",
		FriendId: "cnqowrn42j",
	}
	_, err := cc.AddFriend(context.TODO(), friend)
	if err != nil {
		panic(err)
	}
}

func getFriend() {
	user := &pb.GetUserReq{
		Username: "cnqowrn42j",
	}
	resp, err := cc.GetFriend(context.TODO(), user)
	if err != nil {
		panic(err)
	}
	fmt.Println("[GetFriend]=>", resp.User)
}

func listFriends() {
	user := &pb.ListUsersReq{
		Username: "qffqwrtb231",
	}
	resp, err := cc.ListFriends(context.TODO(), user)
	if err != nil {
		panic(err)
	}
	fmt.Println("[ListFriends]=>", resp.Users)
}

func deleteFriend() {
	del := &pb.DelFriendReq{
		Username:   "qffqwrtb231",
		FriendName: "cnqowrn42j",
	}
	_, err := cc.DelFriend(context.TODO(), del)
	if err != nil {
		panic(err)
	}
}

func updateProfile() {
	req := &pb.UpdateProfileReq{
		User: &pb.User{
			Username: "qffqwrtb231",
			Nickname: "Smith",
			Sex:      1,
			Birthday: time.Date(1995, time.January, 10, 0, 0, 0, 0, time.UTC).Unix(),
			PhoneNum: "12405762905",
		},
	}
	_, err := cc.UpdateProfile(context.TODO(), req)
	if err != nil {
		panic(err)
	}
}

func changePassword() {
	req := &pb.ChangePasswordReq{
		Username:    "qffqwrtb231",
		PrePassword: "123456",
		Password:    "890567",
	}
	_, err := cc.ChangePassword(context.TODO(), req)
	if err != nil {
		panic(err)
	}
}

func sendMessage() {
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
		SenderName:   "cnqowrn42j",
		SenderType:   pb.SenderType_ST_USER,
		ReceiverName: "qffqwrtb231",
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

func tcpConnTest() {
	conn, err := net.Dial("tcp", "127.0.0.1:8090")
	if err != nil {
		panic(err)
	}

	cdc := codec.NewCodec(conn)
	go long_link.ReadResp(cdc, getMessage)

	// 登录
	req := &pb.SignInReq{
		Username: "cnqowrn42j",
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
