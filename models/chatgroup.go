package models

import "gorm.io/gorm"

type ChatGroup struct {
	gorm.Model
	GroupID   int    // 群号
	GroupName string // 群名
}

func (ct *ChatGroup) TableName() string {
	return "chatgroup"
}
