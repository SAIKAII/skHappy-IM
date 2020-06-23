package main

import (
	"fmt"
	"github.com/SAIKAII/skHappy-IM/apis"
	"github.com/SAIKAII/skHappy-IM/infra"
	"github.com/SAIKAII/skHappy-IM/infra/base"
	_ "github.com/SAIKAII/skHappy-IM/internal/logic/service"
	"os"
	"os/signal"
)

func init() {
	b := infra.New()

	infra.Register(&base.DatabaseStarter{})
	infra.Register(&base.RedisStarter{})
	infra.Register(&base.ConnStarter{})
	infra.Register(&base.RPCCliStarter{})

	b.Start()
}

func main() {
	// 注册RPC Server
	go apis.StartCliRPCServer("127.0.0.1:8088")

	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, os.Kill)
	select {
	case <-c:
		fmt.Println("Stop")
	}
}
