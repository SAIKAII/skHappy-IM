package base

import (
	"errors"
	"github.com/SAIKAII/skHappy-IM/cmd/config"
	"github.com/SAIKAII/skHappy-IM/infra"
	"github.com/gomodule/redigo/redis"
	"google.golang.org/grpc"
	_ "google.golang.org/grpc"
	"google.golang.org/grpc/connectivity"
	"google.golang.org/grpc/keepalive"
	"sync"
	"time"
)

type RPCCliStarter struct {
	infra.BaseStarter
}

func (r *RPCCliStarter) Setup(ctx infra.StarterContext) {
	rpcCli = &RPCCli{
		dialOpt: make([]grpc.DialOption, 0),
		clients: make(map[string]*grpc.ClientConn),
	}

	kalv := keepalive.ClientParameters{
		Time:                config.GetDuration("grpc.keepalive.time") * time.Second,    // 每隔10秒ping一次
		Timeout:             config.GetDuration("grpc.keepalive.timeout") * time.Second, // 等待2秒ack，若期间无ack，则该连接被断开
		PermitWithoutStream: true,
	}
	rpcCli.dialOpt = append(rpcCli.dialOpt, grpc.WithKeepaliveParams(kalv))
	rpcCli.dialOpt = append(rpcCli.dialOpt, grpc.WithInsecure())
}

var rpcCli *RPCCli

type RPCCli struct {
	mu      sync.Mutex
	dialOpt []grpc.DialOption
	clients map[string]*grpc.ClientConn
}

func NewRPCCli() *RPCCli {
	return rpcCli
}

func (r *RPCCli) Dialer(key, field string) (*grpc.ClientConn, error) {
	rdConn := RedisConn()
	defer rdConn.Close()
	// 从redis缓存中获取指定用户连接到的服务器IP:Port
	addr, err := redis.String(rdConn.Do("HGET", key, field))
	if err != nil {
		return nil, err
	}

	r.mu.Lock()
	defer r.mu.Unlock()

	var (
		cc *grpc.ClientConn
		ok bool
	)
	if cc, ok = r.clients[addr]; !ok {
		// 之前没有与该服务器进行rpc连接
		var err error
		cc, err = grpc.Dial(addr, r.dialOpt...)
		if err != nil {
			return nil, err
		}

		r.clients[addr] = cc
	}

	state := cc.GetState()
	if state == connectivity.Shutdown {
		cc.Close()
		delete(r.clients, key)
		return nil, errors.New("连接被断开")
	}

	return cc, nil
}
