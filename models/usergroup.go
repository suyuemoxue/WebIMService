package models

import "gorm.io/gorm"

type UserGroup struct {
	gorm.Model
	UserID    string // 用户ID
	GroupID   int    // 群号
	GroupName string // 群名
	Role      string // 群地位 群主 管理员 群员
}

func (ct *UserGroup) TableName() string {
	return "usergroup"
}
