package core

import (
	"WebIM/global"
	"WebIM/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
)

func InitGorm() {
	dsn := global.Config.MySQL.Dsn()
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger:                 logger.Default.LogMode(logger.Silent),
		SkipDefaultTransaction: true,
	})
	if err != nil {
		panic("mysql connect error" + err.Error())
	}
	log.Println("mysql connect success")
	global.DB = db
	global.DB = global.DB.Session(&gorm.Session{})
	err = global.DB.AutoMigrate(&models.User{}) // 初始化用户表
	if err != nil {
		return
	}
	err = global.DB.AutoMigrate(&models.Message{}) // 初始化消息表
	if err != nil {
		return
	}
	err = global.DB.AutoMigrate(&models.Contact{}) // 初始化用户关系表
	if err != nil {
		return
	}
}
