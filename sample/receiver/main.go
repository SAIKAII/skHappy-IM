package main

import (
	"fmt"
	pb "github.com/SAIKAII/skHappy-IM/protocols"
	"github.com/SAIKAII/skHappy-IM/sample/jwt"
	"github.com/SAIKAII/skHappy-IM/sample/long_link"
	codec "github.com/SAIKAII/skHappy-IM/sample/util"
	"google.golang.org/grpc"
	"net"
	"os"
	"os/signal"
)

var cc pb.CliInterfaceServiceClient

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8090")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	c := make(chan interface{})
	defer close(c)
	cdc := codec.NewCodec(conn)
	go long_link.ReadResp(cdc, getMessage, c)

	// 登录
	req := &pb.SignInReq{
		Username: "ituen,vlos1",
		Password: "123456",
	}
	err = long_link.Login(cdc, req)
	if err != nil {
		panic(err)
	}

	go long_link.HeartBeat(cdc)

	jwtString := <-c
	tk := &jwt.JWTSt{JWTString: jwtString.(string)}
	cli, err := grpc.Dial("127.0.0.1:8088", grpc.WithInsecure(), grpc.WithPerRPCCredentials(tk))
	if err != nil {
		panic(err)
	}
	cc = pb.NewCliInterfaceServiceClient(cli)

	// 拉取离线消息
	sync := &pb.SyncReq{
		Username:  "ituen,vlos1",
		LastSeqId: 70,
	}
	err = long_link.Sync(cdc, sync)
	if err != nil {
		panic(err)
	}
	chSig := make(chan os.Signal)
	signal.Notify(chSig, os.Interrupt, os.Kill)
	select {
	case <-chSig:
		cdc.CloseConnection()
		fmt.Println("Stop")
	}
}
