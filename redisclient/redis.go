package redisclient

import (
	"context"
	"github.com/go-redis/redis/v8"
)

type RedisConfig struct {
	Addr     string
	Password string
	DbName   int
}

// 声明一个全局的rdb变量

// 初始化连接
func NewRedisClient(c RedisConfig) (*redis.Client, error) {

	rdb := redis.NewClient(&redis.Options{
		Addr:     c.Addr,
		Password: c.Password, // no password set
		DB:       c.DbName,   // use default DB
	})
	if err := rdb.Ping(context.Background()).Err(); err != nil {
		return nil, err
	}
	return rdb, nil
}
