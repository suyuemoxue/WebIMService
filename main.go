package main

import (
	"WebIM/core"
	"WebIM/routers"
)

func main() {
	core.InitSocket()
	core.InitConfig()    // 读取配置文件config.yaml，并初始化给Config结构体
	core.InitRedis()     // 初始化redis
	core.InitGorm()      // 初始化gorm，连接MySQL
	routers.InitRouter() // 初始化gin，开启服务
}
