package main

import (
	"fmt"
	pb "github.com/SAIKAII/skHappy-IM/protocols"
	codec "github.com/SAIKAII/skHappy-IM/sample/util"
	"github.com/golang/protobuf/proto"
	"log"
	"net"
	"os"
	"os/signal"
	"time"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8090")
	if err != nil {
		panic(err)
	}

	cdc := codec.NewCodec(conn)
	go readResp(cdc)

	// 登录
	err = login(cdc)
	if err != nil {
		panic(err)
	}

	go heartBeat(cdc)

	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, os.Kill)
	select {
	case <-c:
		cdc.CloseConnection()
		fmt.Println("Stop")
	}
}

func readResp(cdc *codec.Codec) {
	for !cdc.IsClosed() {
		n, err := cdc.Read()
		if err != nil {
			log.Println(err)
			break
		} else if n == 0 {
			continue
		}

		b, _, err := cdc.Decode()
		if err != nil {
			log.Println(err)
			break
		}

		var data pb.ConnOutput
		proto.Unmarshal(b, &data)
		switch data.PackageType {
		case pb.PackageType_PT_SIGN_IN:
			fmt.Println("[SignIn]=>", data.ErrCode, data.ErrMsg)
		case pb.PackageType_PT_MESSAGE:
			fmt.Println("[Message]=>", data.ErrCode, data.ErrMsg)
			getMessage(data.Data)
		case pb.PackageType_PT_HEART_BEAT:
			fmt.Println("[HeartBeat]=>", data.ErrCode, data.ErrMsg)
		default:
			fmt.Println("Data Error!")
		}
	}
}

func login(cdc *codec.Codec) error {
	req := &pb.SignInReq{
		Username: "qffqwrtb231",
		Password: "890567",
	}
	o, _ := proto.Marshal(req)
	in := &pb.ConnInput{
		PackageType: pb.PackageType_PT_SIGN_IN,
		Data:        o,
	}
	d, _ := proto.Marshal(in)

	return cdc.Write(cdc.Encode(d))
}

func heartBeat(cdc *codec.Codec) {
	for {
		time.Sleep(1 * time.Second)

		hb := &pb.ConnInput{
			PackageType: pb.PackageType_PT_HEART_BEAT,
			Data:        nil,
		}
		o, _ := proto.Marshal(hb)
		err := cdc.Write(cdc.Encode(o))
		if err != nil {
			log.Println(err)
		}
	}

}

func getMessage(data []byte) error {
	var content pb.MessageOutput
	err := proto.Unmarshal(data, &content)
	if err != nil {
		return err
	}

	fmt.Println(content.Item)
	return nil
}
