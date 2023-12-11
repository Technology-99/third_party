package redisclient

import (
	"context"
	"github.com/go-redis/redis/v8"
)

type RedisConfig struct {
	Host string
	Pass string `json:",optional"`
	Tls  bool   `json:",optional"`
	DB   int    `json:",default=0"`
}

// 声明一个全局的rdb变量

// 初始化连接
func NewRedisClient(c *RedisConfig) (*redis.Client, error) {

	rdb := redis.NewClient(&redis.Options{
		Addr:     c.Host,
		Password: c.Pass, // no password set
		DB:       c.DB,   // use default DB
	})
	if err := rdb.Ping(context.Background()).Err(); err != nil {
		return nil, err
	}
	return rdb, nil
}
