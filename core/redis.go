package core

import (
	"WebIM/global"
	"context"
	"github.com/go-redis/redis/v8"
	"log"
	"time"
)

func InitRedis() {
	global.RDB = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
		PoolSize: 10,
	})
	// 设置超时时间
	timeoutCtx, cancel := context.WithTimeout(global.Ctx, 100*time.Millisecond)
	defer cancel()
	pong, err := global.RDB.Ping(timeoutCtx).Result()
	if err != nil {
		log.Println("Error:", err)
		return
	}
	log.Println("redis connect success,", pong)
}
