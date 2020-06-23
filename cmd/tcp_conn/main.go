package main

import (
	"fmt"
	"github.com/SAIKAII/skHappy-IM/apis"
	"github.com/SAIKAII/skHappy-IM/infra"
	"github.com/SAIKAII/skHappy-IM/infra/base"
	"os"
	"os/signal"
)

func init() {
	b := infra.New()

	infra.Register(&base.DatabaseStarter{})
	infra.Register(&base.RedisStarter{})
	infra.Register(&base.ConnStarter{})

	b.Start()
}

func main() {
	// 注册RPC Server
	go apis.StartConnRPCServer("127.0.0.1:8089")

	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, os.Kill)
	select {
	case <-c:
		fmt.Println("Stop")
	}
}
