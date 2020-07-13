package main

import (
	"fmt"
	coma "github.com/SAIKAII/go-conn-manager"
	"github.com/SAIKAII/skHappy-IM/apis"
	"github.com/SAIKAII/skHappy-IM/infra"
	"github.com/SAIKAII/skHappy-IM/infra/base"
	"github.com/SAIKAII/skHappy-IM/internal/logic/service"
	_ "github.com/SAIKAII/skHappy-IM/internal/logic/service"
	"github.com/spf13/viper"
	"os"
	"os/signal"
	"time"
)

func init() {
	b := infra.New()

	infra.Register(&base.DatabaseStarter{})
	infra.Register(&base.RedisStarter{})
	infra.Register(&base.ConnStarter{})
	infra.Register(&base.LoggerStarter{})

	b.Start()
}

func main() {
	viper.SetConfigName("config")
	viper.AddConfigPath("cmd/config/")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s\n", err))
	}

	// 注册RPC Server
	go apis.StartConnRPCServer(
		fmt.Sprintf("%s:%d",
			viper.GetString("tcp-server.host"),
			viper.GetInt("tcp-server.rpc-port")))

	// 启动TCP监控
	th := service.NewTCPHandler(viper.GetString("tcp-server.host"))
	epoll := coma.NewEpoll(10 * time.Second)
	server := coma.NewServer(epoll)
	go server.Start(
		viper.GetString("tcp-server.host"),
		viper.GetInt("tcp-server.long-conn-port"),
		viper.GetInt("tcp-server.header-len"),
		viper.GetInt("tcp-server.read-max-len"),
		viper.GetInt("tcp-server.write-max-len"),
		th)

	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, os.Kill)
	select {
	case <-c:
		server.Stop()
		fmt.Println("Stop")
	}
}
