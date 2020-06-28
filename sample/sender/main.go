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

	selfName := "cnqowrn42j"
	friendName := "qffqwrtb231"
	// 注册帐号
	//register()
	// 添加好友
	//addReq := &pb.AddFriendReq{
	//	UserId:   selfName,
	//	FriendId: friendName,
	//}
	//addFriend(addReq)
	// 获取好友信息
	getReq := &pb.GetUserReq{
		Username: friendName,
	}
	getFriend(getReq)
	// 列出所有好友
	listReq := &pb.ListUsersReq{
		Username: selfName,
	}
	listFriends(listReq)
	// 删除好友关系
	//delReq := &pb.DelFriendReq{
	//	Username:   selfName,
	//	FriendName: friendName,
	//}
	//deleteFriend(delReq)
	// 更新个人信息
	//updateReq := &pb.UpdateProfileReq{
	//	User: &pb.User{
	//		Username: selfName,
	//		Nickname: "Smith",
	//		Sex:      1,
	//		Birthday: time.Date(1995, time.January, 10, 0, 0, 0, 0, time.UTC).Unix(),
	//		PhoneNum: "12405762905",
	//	},
	//}
	//updateProfile(updateReq)
	// 更换密码
	//newPwdReq := &pb.ChangePasswordReq{
	//	Username:    selfName,
	//	PrePassword: "123456",
	//	Password:    "890567",
	//}
	//changePassword(newPwdReq)
	// 发送消息
	sendMessage(selfName, friendName)

	// tcp测试
	tcpConnTest(selfName)
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

func addFriend(friend *pb.AddFriendReq) {
	_, err := cc.AddFriend(context.TODO(), friend)
	if err != nil {
		panic(err)
	}
}

func getFriend(user *pb.GetUserReq) {
	resp, err := cc.GetFriend(context.TODO(), user)
	if err != nil {
		panic(err)
	}
	fmt.Println("[GetFriend]=>", resp.User)
}

func listFriends(user *pb.ListUsersReq) {
	resp, err := cc.ListFriends(context.TODO(), user)
	if err != nil {
		panic(err)
	}
	fmt.Println("[ListFriends]=>", resp.Users)
}

func deleteFriend(del *pb.DelFriendReq) {
	_, err := cc.DelFriend(context.TODO(), del)
	if err != nil {
		panic(err)
	}
}

func updateProfile(profile *pb.UpdateProfileReq) {
	_, err := cc.UpdateProfile(context.TODO(), profile)
	if err != nil {
		panic(err)
	}
}

func changePassword(newPwd *pb.ChangePasswordReq) {
	_, err := cc.ChangePassword(context.TODO(), newPwd)
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

func tcpConnTest(selfName string) {
	conn, err := net.Dial("tcp", "127.0.0.1:8090")
	if err != nil {
		panic(err)
	}

	cdc := codec.NewCodec(conn)
	go long_link.ReadResp(cdc, getMessage)

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
