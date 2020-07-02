package cache

import (
	"fmt"
	"github.com/SAIKAII/skHappy-IM/infra/base"
	"github.com/SAIKAII/skHappy-IM/internal/logic/dao"
	"github.com/gomodule/redigo/redis"
	jsoniter "github.com/json-iterator/go"
)

const GROUP_KEY = "group_key:"

var GroupUserCache = &groupUserCache{}

type groupUserCache struct {
}

func (g *groupUserCache) Key(groupId uint64) string {
	return fmt.Sprintf("%s:%d", GROUP_KEY, groupId)
}

func (g *groupUserCache) Set(key string, users []*dao.GroupUser) error {
	rdConn := base.RedisConn()
	defer rdConn.Close()

	byt, err := jsoniter.Marshal(users)
	if err != nil {
		return err
	}

	_, err = rdConn.Do("Set", key, byt)
	if err != nil {
		return err
	}

	return nil
}

func (g *groupUserCache) Get(key string) ([]*dao.GroupUser, error) {
	rdConn := base.RedisConn()
	defer rdConn.Close()

	byt, err := redis.Bytes(rdConn.Do("Get", key))
	if err != nil {
		return nil, err
	}

	var users []*dao.GroupUser
	err = jsoniter.Unmarshal(byt, &users)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (g *groupUserCache) Del(key string) error {
	rdConn := base.RedisConn()
	defer rdConn.Close()

	_, err := rdConn.Do("Del", key)
	if err != nil {
		return err
	}

	return nil
}
