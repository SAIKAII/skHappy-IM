package main

import (
	"fmt"
	"github.com/SAIKAII/skHappy-IM/apis"
	"github.com/SAIKAII/skHappy-IM/cmd/config"
	"github.com/SAIKAII/skHappy-IM/infra"
	"github.com/SAIKAII/skHappy-IM/infra/base"
	_ "github.com/SAIKAII/skHappy-IM/internal/logic/service"
	"os"
	"os/signal"
)

var b *infra.BootApplication

func main() {
	config.Init("cmd/config/")
	// 初始化
	Init(infra.StarterContext{})

	// 注册RPC Server
	go apis.StartCliRPCServer(
		fmt.Sprintf("%s:%d",
			config.GetString("cli-rpc-server.host"),
			config.GetInt("cli-rpc-server.port")))

	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, os.Kill)
	select {
	case <-c:
		fmt.Println("Stop")
	}
}

func Init(ctx infra.StarterContext) {
	b = infra.New(ctx)

	infra.Register(&base.DatabaseStarter{})
	infra.Register(&base.RedisStarter{})
	infra.Register(&base.ConnStarter{})
	infra.Register(&base.RPCCliStarter{})
	infra.Register(&base.LoggerStarter{})

	b.Start()
}
