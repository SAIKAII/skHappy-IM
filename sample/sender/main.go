package main

import (
	"fmt"
	pb "github.com/SAIKAII/skHappy-IM/protocols"
	"github.com/SAIKAII/skHappy-IM/sample/jwt"
	"google.golang.org/grpc"
	"os"
	"os/signal"
)

var cc pb.CliInterfaceServiceClient

func main() {
	selfName := RandString(32)
	selfNickname := RandString(32)
	friendName := "ituen,vlos1"
	//friendNickname := "Red"
	tk := &jwt.JWTSt{}
	cli, err := grpc.Dial("127.0.0.1:8088", grpc.WithInsecure(), grpc.WithPerRPCCredentials(tk))
	if err != nil {
		panic(err)
	}
	cc = pb.NewCliInterfaceServiceClient(cli)

	// 注册
	register(selfName, selfNickname)

	// 登录 获取JWT
	c := make(chan interface{}, 1)
	defer close(c)
	// tcp测试
	go tcpConnTest(selfName, c)
	jwtString := <-c
	tk.JWTString = jwtString.(string)

	// 添加好友
	addFriend(selfName, friendName)
	// 获取好友信息
	getFriend(friendName)
	// 获取自己的所有好友
	listFriends(selfName)
	// 更新个人信息
	updateProfile(selfName, "Yoy")
	// 修改密码
	changePassword(selfName, "987654")
	// 删除好友关系
	//deleteFriend(selfName, friendName)
	// 发送消息
	sendMessage(selfName, friendName)

	chSig := make(chan os.Signal)
	signal.Notify(chSig, os.Interrupt, os.Kill)
	select {
	case <-chSig:
		cli.Close()
		fmt.Println("Stop")
	}
}
