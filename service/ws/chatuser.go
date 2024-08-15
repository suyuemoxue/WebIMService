package ws

import (
	"github.com/gorilla/websocket"
)

// SendMsg 发送消息的结构体
type SendMsg struct {
	SendID    string `json:"sendID"`
	ReceiveID string `json:"receiveID"`
	Content   []byte `json:"content"`
	MsgType   string `json:"msgType"`   // 消息类型 群聊 私聊
	MediaType string `json:"mediaType"` // 消息类型 文字 图片 音视频
}

// ReplyMsg 回复消息的结构体
type ReplyMsg struct {
	From string `json:"from"`
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

// Client 用户结构体
type Client struct {
	UserID     string
	Socket     *websocket.Conn
	ReceiveMsg chan []byte
	Online     bool // 在线状态
}

// Group 群组结构体
//type Group struct {
//	GroupID string
//	SendID  string
//	ReceiveMsg chan []byte
//}

// Broadcast 广播类（包括广播内容和源用户）
//type Broadcast struct {
//	Client    *Client
//	Group     *Group
//	Msg   []byte
//	MsgType   string // 消息类型 群聊 私聊
//	MediaType string // 消息类型 文字 图片 音视频
//}

// ClientManager 用户管理
type ClientManager struct {
	Clients map[string]*Client
	//Broadcast  chan *Broadcast
	Reply      chan *Client
	Register   chan *Client
	Unregister chan *Client
	SendMsg    chan *SendMsg
}

// Message 信息转JSON (包括：发送者、接收者、内容)
type Message struct {
	Sender    string `json:"sender,omitempty"`
	Recipient string `json:"recipient,omitempty"`
	Content   string `json:"content,omitempty"`
}

// Manager 定义一个管理Manager
var Manager = ClientManager{
	Clients: make(map[string]*Client), // 参与连接的用户，出于性能的考虑，需要设置最大连接数
	//Broadcast:  make(chan *Broadcast),
	SendMsg:    make(chan *SendMsg),
	Register:   make(chan *Client),
	Reply:      make(chan *Client),
	Unregister: make(chan *Client),
}
