package service

import (
	coma "github.com/SAIKAII/go-conn-manager"
	"github.com/SAIKAII/skHappy-IM/infra/base"
	pb "github.com/SAIKAII/skHappy-IM/protocols"
	"github.com/SAIKAII/skHappy-IM/services"
	"github.com/SAIKAII/skHappy-IM/services/common"
	"github.com/golang/protobuf/proto"
	"github.com/gomodule/redigo/redis"
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
	}
}

// OnClose 主动关闭连接或超时无心跳包时调用
func (th *TCPHandler) OnClose(conn *coma.Conn) error {
	cData := conn.Data().(*ConnData)
	rdConn := base.RedisConn()
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
		resp.ErrCode = common.CLIENT_REQUEST_PARAMS_ERROR
		resp.ErrMsg = err.Error()
		o, _ := proto.Marshal(resp)
		return coma.PacketToPeer(conn, o)
	}

	// 验证用户登录是否通过
	err = services.IAuthService.SignInAuth(input.Username, input.Password)
	if err != nil {
		// 验证没通过
		resp.ErrMsg = err.Error()
		if err == services.AUTH_FAILURE {
			resp.ErrCode = common.COMMON_PWD_NOT_MATCH_ERROR
		} else if err == services.USER_NOT_FOUND {
			resp.ErrCode = common.COMMON_USER_NOT_FOUND_ERROR
		} else {
			resp.ErrCode = common.COMMON_UNKNOWN_ERROR
		}
		o, _ := proto.Marshal(resp)
		return coma.PacketToPeer(conn, o)
	}

	rdConn := base.RedisConn()
	_, err = redis.Bool(rdConn.Do("HSET", base.USER_ADDR, input.Username, th.host))
	if err != nil {
		resp.ErrCode = common.INTERNEL_UNKNOWN_ERROR
		resp.ErrMsg = "保存登录状态失败"
		o, _ := proto.Marshal(resp)
		return coma.PacketToPeer(conn, o)
	}

	base.ConnectionManager().StoreConn(input.Username, conn)
	// 保存一些与该连接有关的信息
	cData := &ConnData{Username: input.Username}
	conn.SetData(cData)

	output := &pb.SignInResp{}
	b, _ := proto.Marshal(output)
	resp.Data = b
	o, _ := proto.Marshal(resp)

	return coma.PacketToPeer(conn, o)
}

func (th *TCPHandler) sync(conn *coma.Conn, data []byte) error {
	return nil
}

func (th *TCPHandler) heartBeat(conn *coma.Conn, data []byte) error {
	cData := conn.Data().(*ConnData)
	rdConn := base.RedisConn()
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
