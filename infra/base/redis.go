package base

import (
	"github.com/SAIKAII/skHappy-IM/infra"
	"github.com/gomodule/redigo/redis"
)

var redisDB *redis.Pool
var addr = "127.0.0.1:6379"

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
	redisDB = &redis.Pool{
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", addr)
		},
		TestOnBorrow:    nil,
		MaxIdle:         3,
		MaxActive:       20,
		IdleTimeout:     0,
		Wait:            false,
		MaxConnLifetime: 0,
	}
	// 测试是否可以创建连接
	c := redisDB.Get()
	defer c.Close()
	_, err := c.Do("PING")
	if err != nil {
		panic(err)
	}
}
