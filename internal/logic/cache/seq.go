package cache

import (
	"github.com/SAIKAII/skHappy-IM/infra/base"
	"github.com/gomodule/redigo/redis"
)

const USER_SEQ = "user_seq:"

func UserKey(username string) string {
	return USER_SEQ + username
}

func UpdateUserSeq(name string, seqId uint64) error {
	rdConn := base.RedisConn()
	defer rdConn.Close()
	_, err := rdConn.Do("SET", name, seqId)
	if err != nil {
		return err
	}

	return nil
}

func Incr(name string) (uint64, error) {
	rdConn := base.RedisConn()
	defer rdConn.Close()
	seqId, err := redis.Uint64(rdConn.Do("INCR", name))
	if err != nil {
		return 0, err
	}

	return seqId, nil
}

func Decr(name string) (uint64, error) {
	rdConn := base.RedisConn()
	defer rdConn.Close()
	seqId, err := redis.Uint64(rdConn.Do("DECR", name))
	if err != nil {
		return 0, err
	}

	return seqId, nil
}
