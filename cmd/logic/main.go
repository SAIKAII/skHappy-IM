package main

import (
	"fmt"
	"github.com/SAIKAII/skHappy-IM/apis"
	"github.com/SAIKAII/skHappy-IM/infra"
	"github.com/SAIKAII/skHappy-IM/infra/base"
	_ "github.com/SAIKAII/skHappy-IM/internal/logic/service"
	"github.com/spf13/viper"
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
	viper.SetConfigName("config")
	viper.AddConfigPath("cmd/config/")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s\n", err))
	}

	// 注册RPC Server
	go apis.StartCliRPCServer(
		fmt.Sprintf("%s:%d",
			viper.GetString("cli-rpc-server.host"),
			viper.GetInt("cli-rpc-server.port")))

	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, os.Kill)
	select {
	case <-c:
		fmt.Println("Stop")
	}
}
