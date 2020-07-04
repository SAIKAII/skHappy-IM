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
	//username := util.RandString(32)
	//nickname := util.RandString(32)

	tk := &jwt.JWTSt{}
	cli, err := grpc.Dial("127.0.0.1:8088", grpc.WithInsecure(), grpc.WithPerRPCCredentials(tk))
	if err != nil {
		panic(err)
	}
	defer cli.Close()
	cc = pb.NewCliInterfaceServiceClient(cli)

	// 注册
	//err = register(username, nickname)
	//if err != nil {
	//	panic(err)
	//}

	// 登录 获取JWT
	c := make(chan interface{}, 1)
	defer close(c)

	go tcpConnTest("XzhKSrcVoauCNBkkbVuQFnhaLVQxLvsW", c)
	jwtString := <-c
	tk.JWTString = jwtString.(string)

	// 创建群组
	//_, err = createGroup(username)
	//if err != nil {
	//	panic(err)
	//}

	// 发送群消息
	for i := 0; i < 5; i++ {
		err = sendGroupMessage("XzhKSrcVoauCNBkkbVuQFnhaLVQxLvsW", 12)
		if err != nil {
			panic(err)
		}
	}

	chSig := make(chan os.Signal)
	signal.Notify(chSig, os.Interrupt, os.Kill)
	select {
	case <-chSig:
		fmt.Println("Stop")
	}
}
