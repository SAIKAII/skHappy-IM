package base

import (
	"fmt"
	"github.com/SAIKAII/skHappy-IM/cmd/config"
	"github.com/SAIKAII/skHappy-IM/infra"
	"github.com/gomodule/redigo/redis"
)

var redisDB *redis.Pool

const (
	USER_ADDR = "user_addr" // 用户长连接到的服务器地址
)

func RedisConn() redis.Conn {
	return redisDB.Get()
}

type RedisStarter struct {
	infra.BaseStarter
}

func (r *RedisStarter) Setup(ctx infra.StarterContext) {
	host := config.GetString("redis.host")
	port := config.GetInt("redis.port")
	redisDB = &redis.Pool{
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", fmt.Sprintf("%s:%d", host, port))
		},
		TestOnBorrow:    nil,
		MaxIdle:         config.GetInt("redis.max-idle"),
		MaxActive:       config.GetInt("redis.max-active"),
		IdleTimeout:     config.GetDuration("redis.idle-timeout"),
		Wait:            false,
		MaxConnLifetime: config.GetDuration("redis.max-conn-lifetime"),
	}
	// 测试是否可以创建连接
	c := redisDB.Get()
	defer c.Close()
	_, err := c.Do("PING")
	if err != nil {
		panic(err)
	}
}
