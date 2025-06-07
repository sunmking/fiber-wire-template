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

			if setting.RedisCnf.Password != "" { // Corrected from Pass to Password
				if _, err := c.Do("AUTH", setting.RedisCnf.Password); err != nil {
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
			// Panicking on a connection close error might be too aggressive.
			// Consider logging the error instead, as the pool handles connection health.
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
			// Panicking on a connection close error might be too aggressive.
			// Consider logging the error instead, as the pool handles connection health.
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
			// Panicking on a connection close error might be too aggressive.
			// Consider logging the error instead, as the pool handles connection health.
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
			// Panicking on a connection close error might be too aggressive.
			// Consider logging the error instead, as the pool handles connection health.
			panic(fmt.Sprintf("redis error: %s", err.Error()))
		}
	}(conn)
	return redis.Bool(conn.Do("DEL", key))
}

// LikeDeletes WARNING: Using KEYS in production Redis is dangerous and can block the server.
// This command may cause performance issues for Redis instances with a large number of keys.
// Consider using Redis sets to manage collections of keys for bulk deletion,
// or use SCAN for iterative deletion if pattern matching is absolutely necessary.
// SCAN is less blocking but still requires careful implementation for bulk deletes.
// For safety, this function will not execute if the key is empty or "*".
func (r *Redis) LikeDeletes(key string) error {
	if key == "" || key == "*" {
		return fmt.Errorf("redis: unsafe LikeDeletes pattern '%s'. Aborting KEYS command", key)
	}

	conn := r.RedisConn.Get()
	defer func(conn redis.Conn) {
		err := conn.Close()
		if err != nil {
			// Panicking on a connection close error might be too aggressive.
			// Consider logging the error instead, as the pool handles connection health.
			panic(fmt.Sprintf("redis error: %s", err.Error()))
		}
	}(conn)

	// Construct the pattern carefully. The original code used "*"+key+"*".
	// Ensure this is the intended pattern. For example, if key is "user:123", pattern becomes "*user:123*".
	pattern := "*" + key + "*"
	keys, err := redis.Strings(conn.Do("KEYS", pattern))
	if err != nil {
		return err
	}

	if len(keys) == 0 {
		return nil // No keys matched, nothing to delete.
	}

	// Deleting keys one by one can be slow.
	// Consider using a pipeline or DEL with multiple arguments if appropriate and supported.
	// However, since Delete() itself gets a new connection, pipelining here is complex.
	// For a large number of keys, this loop can be very slow.
	for _, k := range keys { // Renamed loop variable to avoid shadowing 'key' parameter
		_, err = r.Delete(k) // Use the loop variable 'k'
		if err != nil {
			// Decide on error handling: return immediately or collect errors?
			// Returning immediately might leave some keys undeleted.
			return fmt.Errorf("redis: error deleting key '%s': %w", k, err)
		}
	}
	return nil
}
