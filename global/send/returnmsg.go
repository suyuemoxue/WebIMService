package send

type ReturnMsg struct {
	SendID    string `json:"sendId"`
	ReceiveID string `json:"receiveId"`
	Content   string `json:"content"`
	MsgType   string `json:"msgType"`   // 消息类型 群聊 私聊
	MediaType string `json:"mediaType"` // 消息类型 文字 图片 音视频
}

type ReturnMsgList []ReturnMsg
