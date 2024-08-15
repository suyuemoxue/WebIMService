package models

import "gorm.io/gorm"

type Contact struct {
	gorm.Model
	OwnerID  string // 己方ID
	FriendID string // 好友ID
}

func (ct *Contact) TableName() string {
	return "contact"
}
