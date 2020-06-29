package service

import (
	"errors"
	"fmt"
	coma "github.com/SAIKAII/go-conn-manager"
	"github.com/SAIKAII/skHappy-IM/infra/base"
	"github.com/SAIKAII/skHappy-IM/internal/logic/dao"
	"github.com/SAIKAII/skHappy-IM/pkg/jwt"
	pb "github.com/SAIKAII/skHappy-IM/protocols"
	"github.com/SAIKAII/skHappy-IM/services"
	"github.com/SAIKAII/skHappy-IM/services/common"
	"github.com/golang/protobuf/proto"
	"github.com/gomodule/redigo/redis"
	"log"
)

type ConnData struct {
	Username string
}

type TCPHandler struct {
	host string
}

// NewTCPHandler 创建TCPHandler实例
func NewTCPHandler(host string) *TCPHandler {
	return &TCPHandler{host: host}
}

// OnConnect 创建连接时调用
func (th *TCPHandler) OnConnect(*coma.Conn) {

}

// OnMessage 套接字有消息可读时调用
func (th *TCPHandler) OnMessage(conn *coma.Conn, data []byte) {
	var input pb.ConnInput
	err := proto.Unmarshal(data, &input)
	if err != nil {

	}

	switch input.PackageType {
	case pb.PackageType_PT_SIGN_IN:
		err = th.signIn(conn, input.Data)
	case pb.PackageType_PT_SYNC_MESSAGE:
		err = th.sync(conn, input.Data)
	case pb.PackageType_PT_HEART_BEAT:
		err = th.heartBeat(conn, input.Data)
	case pb.PackageType_PT_MESSAGE:
		err = th.messageAck(conn, input.Data)
	default:

	}
	if err != nil {
		// TODO
		log.Println(err)
	}
}

// OnClose 主动关闭连接或超时无心跳包时调用
func (th *TCPHandler) OnClose(conn *coma.Conn) error {
	tmp := conn.Data()
	if tmp == nil {
		return errors.New("无该用户的相关连接信息")
	}
	cData := tmp.(*ConnData)
	rdConn := base.RedisConn()
	defer rdConn.Close()
	_, err := rdConn.Do("HDEL", base.USER_ADDR, cData.Username)
	base.ConnectionManager().DeleteConn(cData.Username)
	return err
}

// OnError 套接字发生了错误，一般是接收到RST
func (th *TCPHandler) OnError(conn *coma.Conn) {
	th.OnClose(conn)
}

func (th *TCPHandler) signIn(conn *coma.Conn, data []byte) error {
	var input pb.SignInReq
	resp := &pb.ConnOutput{
		PackageType: pb.PackageType_PT_SIGN_IN,
		Data:        nil,
	}
	err := proto.Unmarshal(data, &input)
	if err != nil {
		// 对请求的数据unmarshal失败
		return th.handleError(conn, resp, common.CLIENT_REQUEST_PARAMS_ERROR, err.Error())
	}

	// 验证用户登录是否通过
	err = services.IAuthService.SignInAuth(input.Username, input.Password)
	if err != nil {
		// 验证没通过
		if err == services.AUTH_FAILURE {
			return th.handleError(conn, resp, common.COMMON_PWD_NOT_MATCH_ERROR, err.Error())
		} else if err == services.USER_NOT_FOUND {
			return th.handleError(conn, resp, common.COMMON_USER_NOT_FOUND_ERROR, err.Error())
		} else {
			return th.handleError(conn, resp, common.COMMON_UNKNOWN_ERROR, err.Error())
		}
	}

	// 生成JWT
	meta := make(map[string]interface{})
	meta["username"] = input.Username
	jwtString, err := jwt.NewJWT(meta)
	if err != nil {
		return th.handleError(conn, resp, common.INTERNEL_GENERATE_JWT_ERROR, "JWT生成失败")
	}

	output := &pb.SignInResp{Jwt: jwtString}

	rdConn := base.RedisConn()
	defer rdConn.Close()
	_, err = redis.Bool(rdConn.Do("HSET", base.USER_ADDR, input.Username, fmt.Sprintf("%s:%d", th.host, 8089)))
	if err != nil {
		return th.handleError(conn, resp, common.INTERNEL_UNKNOWN_ERROR, "保存登录状态失败")
	}

	base.ConnectionManager().StoreConn(input.Username, conn)
	// 保存一些与该连接有关的信息
	cData := &ConnData{Username: input.Username}
	conn.SetData(cData)

	b, _ := proto.Marshal(output)
	resp.Data = b
	o, _ := proto.Marshal(resp)

	return coma.PacketToPeer(conn, o)
}

func (th *TCPHandler) sync(conn *coma.Conn, data []byte) error {
	var input pb.SyncReq
	output := &pb.ConnOutput{
		PackageType: pb.PackageType_PT_SYNC_MESSAGE,
		Data:        nil,
	}
	err := proto.Unmarshal(data, &input)
	if err != nil {
		return th.handleError(conn, output, common.CLIENT_REQUEST_PARAMS_ERROR, err.Error())
	}

	db := base.Database()
	msgDao := dao.MessageDao{DB: db}
	res, err := msgDao.GetAllRecvByLastSeqId(input.Username, input.LastSeqId)
	if err != nil {
		return th.handleError(conn, output, common.INTERNEL_UNKNOWN_ERROR, err.Error())
	}

	var resp pb.SyncResp
	resp.Msg = make([]*pb.MessageItem, 0, len(res))
	resp.HasMore = false
	for _, v := range res {
		mb := &pb.MessageBody{
			Type: pb.MessageType(v.Type),
		}
		if pb.MessageType(v.Type) == pb.MessageType_MT_TEXT {
			t := &pb.Text{Text: v.Content}
			ct := &pb.MessageContent_Text{Text: t}
			mb.Content = &pb.MessageContent{Content: ct}
		} else if pb.MessageType(v.Type) == pb.MessageType_MT_IMAGE {
			// TODO
		}
		item := &pb.MessageItem{
			SenderName:   v.Sender,
			SenderType:   pb.SenderType(v.SenderType),
			ReceiverName: v.Receiver,
			ReceiverType: pb.ReceiverType(v.ReceiverType),
			MsgBody:      mb,
			SendTime:     v.SendTime.Unix(),
		}
		resp.Msg = append(resp.Msg, item)
	}
	d, _ := proto.Marshal(&resp)
	output.Data = d
	o, _ := proto.Marshal(output)

	return coma.PacketToPeer(conn, o)
}

func (th *TCPHandler) heartBeat(conn *coma.Conn, data []byte) error {
	tmp := conn.Data()
	if tmp == nil {
		return errors.New("无该用户的相关连接信息")
	}
	cData := tmp.(*ConnData)
	rdConn := base.RedisConn()
	defer rdConn.Close()
	exist, err := redis.Bool(rdConn.Do("HEXISTS", base.USER_ADDR, cData.Username))
	if err != nil {
		return err
	}

	// 如果缓存中不存在该连接的信息，则重新保存
	if !exist {
		ok, err := redis.Bool(rdConn.Do("HSET", base.USER_ADDR, cData.Username, th.host))
		if err != nil || !ok {
			return err
		}
	}

	return nil
}

func (th *TCPHandler) messageAck(conn *coma.Conn, data []byte) error {
	return nil
}

func (th *TCPHandler) handleError(conn *coma.Conn, resp *pb.ConnOutput, errCode int32, errMsg string) error {
	resp.ErrCode = errCode
	resp.ErrMsg = errMsg
	o, _ := proto.Marshal(resp)
	return coma.PacketToPeer(conn, o)
}
