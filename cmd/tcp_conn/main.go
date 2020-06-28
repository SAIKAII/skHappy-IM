package main

import (
	"fmt"
	coma "github.com/SAIKAII/go-conn-manager"
	"github.com/SAIKAII/skHappy-IM/apis"
	"github.com/SAIKAII/skHappy-IM/infra"
	"github.com/SAIKAII/skHappy-IM/infra/base"
	"github.com/SAIKAII/skHappy-IM/internal/logic/service"
	_ "github.com/SAIKAII/skHappy-IM/internal/logic/service"
	"os"
	"os/signal"
	"time"
)

func init() {
	b := infra.New()

	infra.Register(&base.DatabaseStarter{})
	infra.Register(&base.RedisStarter{})
	infra.Register(&base.ConnStarter{})

	b.Start()
}

func main() {
	tcpAddr := "127.0.0.1"
	// 注册RPC Server
	go apis.StartConnRPCServer(fmt.Sprintf("%s:%d", tcpAddr, 8089))

	// 启动TCP监控
	th := service.NewTCPHandler(tcpAddr)
	epoll := coma.NewEpoll(10 * time.Second)
	server := coma.NewServer(epoll)
	go server.Start("127.0.0.1", 8090, 2, 512, 512, th)

	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, os.Kill)
	select {
	case <-c:
		server.Stop()
		fmt.Println("Stop")
	}
}
