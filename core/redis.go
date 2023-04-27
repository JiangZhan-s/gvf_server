package core

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/sirupsen/logrus"
	"gvf_server/global"
	"time"
)

func ConnectRedis() *redis.Client {
	return ConnectRedisDB(0)
}

func ConnectRedisDB(db int) *redis.Client {
	redisConf := global.Config.Redis
	rdb := redis.NewClient(&redis.Options{
		Addr:     redisConf.Addr(),
		Password: redisConf.Password,
		DB:       db,
		PoolSize: redisConf.PoolSize,
	})
	_, cancel := context.WithTimeout(context.Background(), 500*time.Microsecond)
	defer cancel()
	_, err := rdb.Ping(context.Background()).Result()
	if err != nil {
		logrus.Error("redis连接失败 %s", redisConf.Addr())
		return nil
	}
	return rdb
}
