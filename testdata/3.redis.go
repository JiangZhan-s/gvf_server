package main

//
//import (
//	"context"
//	"fmt"
//	"github.com/go-redis/redis"
//	"github.com/sirupsen/logrus"
//	"time"
//)
//
//var rdb *redis.Client
//
//func init() {
//	rdb = redis.NewClient(&redis.Options{
//		Addr:     "127.0.0.1:6379",
//		Password: "redis",
//		DB:       0,
//		PoolSize: 100,
//	})
//	_, cancel := context.WithTimeout(context.Background(), 500*time.Microsecond)
//	defer cancel()
//	_, err := rdb.Ping(context.Background()).Result()
//	if err != nil {
//		logrus.Error(err)
//		return
//	}
//}
//
//func main() {
//	err := rdb.Set(context.Background(), "xxx1", "value1", 10*time.Second).Err()
//	fmt.Println(err)
//}
