package global

import (
	"WebIM/config"
	"context"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

// 声明一系列全局变量
var (
	Config *config.Config // 用于保存配置文件
	DB     *gorm.DB       // 连接mysql数据库
	RDB    *redis.Client
	Ctx    = context.Background()
)

// Response 用于返回请求的的结构体
type Response struct {
	Code int    `json:"code"`
	Data any    `json:"data"`
	Msg  string `json:"msg"`
}
