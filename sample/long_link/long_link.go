package long_link

import (
	"fmt"
	pb "github.com/SAIKAII/skHappy-IM/protocols"
	codec "github.com/SAIKAII/skHappy-IM/sample/util"
	"github.com/golang/protobuf/proto"
	"log"
	"time"
)

type MessageHandle func([]byte) error

func ReadResp(cdc *codec.Codec, mh MessageHandle) {
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
			mh(data.Data)
		case pb.PackageType_PT_HEART_BEAT:
			fmt.Println("[HeartBeat]=>", data.ErrCode, data.ErrMsg)
		case pb.PackageType_PT_SYNC_MESSAGE:

		default:
			fmt.Println("Data Error!")
		}
	}
	fmt.Println("[ReadResp] Closed")
}

func Login(cdc *codec.Codec, req *pb.SignInReq) error {
	o, _ := proto.Marshal(req)
	in := &pb.ConnInput{
		PackageType: pb.PackageType_PT_SIGN_IN,
		Data:        o,
	}
	d, _ := proto.Marshal(in)

	return cdc.Write(cdc.Encode(d))
}

func HeartBeat(cdc *codec.Codec) {
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
