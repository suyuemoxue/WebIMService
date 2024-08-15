package models

import (
	"WebIM/global"
	"gorm.io/gorm"
)

type Message struct {
	gorm.Model
	SendID    string `json:"sendID" gorm:"size:10" required:"true"`          // 发送者
	ReceiveID string `json:"receiveID" gorm:"size:10" required:"true"`       // 接受者
	Content   []byte `json:"content" gorm:"charset=utf8mb4" required:"true"` // 消息
	MsgType   string `json:"msgType" gorm:"size:10" required:"true"`         // 消息类型 群聊 私聊
	MediaType string `json:"mediaType" gorm:"size:10" required:"true"`       // 消息类型 文字 图片 音视频
	//IsRead    int    `json:"isRead" gorm:"size:1" required:"true"`           // 是否已读,0表示未读，1表示已读
}

type MessageList []Message

func (msg *Message) TableName() string {
	return "message"
}

// SaveMessage 保存消息
func (msg *Message) SaveMessage() bool {
	err := global.DB.Debug().Create(&msg).Error
	if err != nil {
		return false
	}
	return true
}

// GetHistoryMsg 获取历史消息
func (msg *Message) GetHistoryMsg(userName, targetName string) (msgList MessageList, count int64) {
	msgList = make(MessageList, 10)
	count = global.DB.Debug().Where("(send_id = ? AND receive_id = ?) OR (send_id = ? AND receive_id = ?)", userName, targetName, targetName, userName).Find(&msgList).RowsAffected
	return
}
