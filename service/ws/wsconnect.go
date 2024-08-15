package ws

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"sync"
)

type WebsocketService struct {
}

var userLock sync.RWMutex

// WebSocket连接池
var websocketPool = sync.Pool{
	New: func() interface{} {
		return &websocket.Conn{}
	},
}

func (ws *WebsocketService) HandleWebSocket(c *gin.Context) {
	//获取当前用户信息
	query := c.Request.URL.Query()
	userID := query.Get("uid")

	// 处理客户端发起的 WebSocket 连接升级请求的逻辑，如果升级过程顺利完成，那么后续就可以通过连接对象（conn）进行 WebSocket 消息的收发操作
	conn, err := (&websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}).Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println("ws connect fail")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		conn.Close()
		return
	}
	log.Println("用户:", userID, "ws connect success")
	// 从连接池获取连接
	wsConn := websocketPool.Get().(*websocket.Conn)
	*wsConn = *conn

	client := &Client{
		UserID:     userID,
		Socket:     conn,
		ReceiveMsg: make(chan []byte),
	}
	Manager.Register <- client
	go client.Read()
	//go client.SubscribeRedis(global.Ctx, userID)
	//go client.Write()
}
