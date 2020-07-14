package main

import (
	"fmt"
	coma "github.com/SAIKAII/go-conn-manager"
	"github.com/SAIKAII/skHappy-IM/apis"
	"github.com/SAIKAII/skHappy-IM/cmd/config"
	"github.com/SAIKAII/skHappy-IM/infra"
	"github.com/SAIKAII/skHappy-IM/infra/base"
	"github.com/SAIKAII/skHappy-IM/internal/logic/service"
	_ "github.com/SAIKAII/skHappy-IM/internal/logic/service"
	"github.com/spf13/viper"
	"os"
	"os/signal"
	"time"
)

func main() {
	config.Init("cmd/config/")

	// 初始化
	Init(infra.StarterContext{})

	// 注册RPC Server
	go apis.StartConnRPCServer(
		fmt.Sprintf("%s:%d",
			config.GetString("tcp-server.host"),
			config.GetInt("tcp-server.rpc-port")))

	// 启动TCP监控
	th := service.NewTCPHandler(viper.GetString("tcp-server.host"))
	epoll := coma.NewEpoll(10 * time.Second)
	server := coma.NewServer(epoll)
	go server.Start(
		config.GetString("tcp-server.host"),
		config.GetInt("tcp-server.long-conn-port"),
		config.GetInt("tcp-server.header-len"),
		config.GetInt("tcp-server.read-max-len"),
		config.GetInt("tcp-server.write-max-len"),
		th)

	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, os.Kill)
	select {
	case <-c:
		server.Stop()
		fmt.Println("Stop")
	}
}

func Init(ctx infra.StarterContext) {
	b := infra.New(ctx)

	infra.Register(&base.DatabaseStarter{})
	infra.Register(&base.RedisStarter{})
	infra.Register(&base.ConnStarter{})
	infra.Register(&base.LoggerStarter{})

	b.Start()
}
