package gredis

import (
	"encoding/json"
	"fiber-wire-template/pkg/config"
	"fmt"
	"time"

	"github.com/gomodule/redigo/redis"
)

type Redis struct {
	RedisConn *redis.Pool
}

func NewRedis(setting *config.Config) *Redis {
	RedisConn := &redis.Pool{
		MaxIdle:     setting.RedisCnf.MaxIdle,
		MaxActive:   setting.RedisCnf.MaxActive,
		IdleTimeout: time.Duration(setting.RedisCnf.IdleTimeout),
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", setting.RedisCnf.Host)
			if err != nil {
				return nil, err
			}

			if setting.RedisCnf.Pass != "" {
				if _, err := c.Do("AUTH", setting.RedisCnf.Pass); err != nil {
					err := c.Close()
					if err != nil {
						return nil, err
					}
					return nil, err
				}
			}
			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}
	return &Redis{
		RedisConn: RedisConn,
	}
}

func (r *Redis) Set(key string, data interface{}, time int) error {
	conn := r.RedisConn.Get()
	defer func(conn redis.Conn) {
		err := conn.Close()
		if err != nil {
			panic(fmt.Sprintf("redis error: %s", err.Error()))
		}
	}(conn)

	value, err := json.Marshal(data)
	if err != nil {
		return err
	}

	_, err = conn.Do("SET", key, value)
	if err != nil {
		return err
	}

	_, err = conn.Do("EXPIRE", key, time)

	if err != nil {
		return err
	}

	return nil
}

func (r *Redis) Exists(key string) bool {
	conn := r.RedisConn.Get()
	defer func(conn redis.Conn) {
		err := conn.Close()
		if err != nil {
			panic(fmt.Sprintf("redis error: %s", err.Error()))
		}
	}(conn)
	exists, err := redis.Bool(conn.Do("EXISTS", key))
	if err != nil {
		return false
	}
	return exists
}

func (r *Redis) Get(key string) ([]byte, error) {
	conn := r.RedisConn.Get()
	defer func(conn redis.Conn) {
		err := conn.Close()
		if err != nil {
			panic(fmt.Sprintf("redis error: %s", err.Error()))
		}
	}(conn)
	reply, err := redis.Bytes(conn.Do("GET", key))
	if err != nil {
		return nil, err
	}
	return reply, nil
}

func (r *Redis) Delete(key string) (bool, error) {
	conn := r.RedisConn.Get()
	defer func(conn redis.Conn) {
		err := conn.Close()
		if err != nil {
			panic(fmt.Sprintf("redis error: %s", err.Error()))
		}
	}(conn)
	return redis.Bool(conn.Do("DEL", key))
}

func (r *Redis) LikeDeletes(key string) error {
	conn := r.RedisConn.Get()
	defer func(conn redis.Conn) {
		err := conn.Close()
		if err != nil {
			panic(fmt.Sprintf("redis error: %s", err.Error()))
		}
	}(conn)
	keys, err := redis.Strings(conn.Do("KEYS", "*"+key+"*"))
	if err != nil {
		return err
	}
	for _, key := range keys {
		_, err = r.Delete(key)
		if err != nil {
			return err
		}
	}
	return nil
}
