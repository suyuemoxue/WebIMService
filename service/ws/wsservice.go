package ws

import (
	"WebIM/global"
	"WebIM/models"
	"context"
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"log"
	"time"
)

func (cm *ClientManager) Start() {
	for {
		log.Println("----------监听管道通信----------")
		select {
		case conn := <-Manager.Register: // 监听用户连接
			fmt.Printf("有新的连接: %v\n", conn.UserID)
			conn.Online = true
			Manager.Clients[conn.UserID] = conn
			replyMsg := ReplyMsg{
				From: "系统消息",
				Code: 50002,
				Msg:  "已经连接到服务器",
			}
			msg, _ := json.Marshal(replyMsg)
			_ = conn.Socket.WriteMessage(1, msg)
			// 检查是否有存储的离线消息
			err := conn.SendOfflineMessages()
			if err != nil {
				logrus.Error(err)
			}
			go conn.SubscribeRedis(global.Ctx, conn.UserID)
		case conn := <-Manager.Unregister: // 监听用户断开连接
			fmt.Printf("连接断开%s\n", conn.UserID)
			conn.Online = false
			if _, ok := Manager.Clients[conn.UserID]; ok {
				replyMsg := &ReplyMsg{
					From: "系统消息",
					Code: 50003,
					Msg:  "连接中断",
				}
				msg, _ := json.Marshal(replyMsg)
				_ = conn.Socket.WriteMessage(1, msg)
				//close(conn.ReceiveMsg)
				//delete(Manager.Clients, conn.UserID)
			}
		}
	}
}

// 从客户端ws不断读取消息(json格式)，并将消息放进SendMsg{}
func (c *Client) Read() {
	defer func() {
		Manager.Unregister <- c
		_ = c.Socket.Close()
	}()
	for {
		var data map[string]any
		err := c.Socket.ReadJSON(&data) // 从客户端中读取一个JSON格式的消息到data中
		fmt.Println(data)
		if err != nil {
			logrus.Error(err)
			Manager.Unregister <- c
			_ = c.Socket.Close()
			break
		}
		// 验证字段是否为空
		if data["sendId"] == "" || data["receiveId"] == "" || data["message"] == "" || data["msgType"] == "" || data["mediaType"] == "" {
			msg := "有字段为空"
			image, _ := json.Marshal(msg)
			_ = c.Socket.WriteMessage(1, image)
			continue // 如果有空字段，则继续循环等待下一条消息
		}
		// 解析消息类型
		sendMsg := SendMsg{
			SendID:    data["sendId"].(string),
			ReceiveID: data["receiveId"].(string),
			Content:   []byte(data["content"].(string)),
			MsgType:   data["msgType"].(string),
			MediaType: data["mediaType"].(string),
		}
		// 保存消息
		msg := &models.Message{
			SendID:    sendMsg.SendID,
			ReceiveID: sendMsg.ReceiveID,
			Content:   sendMsg.Content,
			MsgType:   sendMsg.MsgType,
			MediaType: sendMsg.MediaType,
			//IsRead:    0,
		}
		res := msg.SaveMessage()
		if !res {
			fmt.Println("保存失败")
			return
		}
		fmt.Println("保存成功")
		// 检查接收方是否在线
		receiverClient, ok := Manager.Clients[sendMsg.ReceiveID]
		// 接收方不在线，将消息存储到 Redis 队列中
		if (Manager.Clients[sendMsg.ReceiveID]) == nil || !receiverClient.Online || !ok {
			if !ok || !receiverClient.Online {
				err := c.StoreMessageToQueue(data)
				if err != nil {
					logrus.Error(err)
				}
			}
		}
		go c.PublishRedis(global.Ctx, sendMsg.ReceiveID, data) // 推送到redis
	}
}

// StoreMessageToQueue 将消息存储到 Redis 队列中
func (c *Client) StoreMessageToQueue(msg any) error {
	m := msg.(map[string]any)
	// 将消息存储到 Redis 队列中
	msgBytes, err := json.Marshal(msg)
	if err != nil {
		return err
	}
	// 这里使用 Redis 的 List 数据结构存储消息
	_, err = global.RDB.RPush(global.Ctx, "offline_messages:"+m["receiveId"].(string), string(msgBytes)).Result()
	if err != nil {
		return err
	}
	return nil
}

// SendOfflineMessages 从 Redis 队列中获取离线消息
func (c *Client) SendOfflineMessages() error {
	// 从 Redis 队列中获取离线消息,然后保存在messages
	messages, err := global.RDB.LRange(global.Ctx, "offline_messages:"+c.UserID, 0, -1).Result()
	if err != nil {
		return err
	}
	// 遍历离线消息，并发送给用户
	for _, msgStr := range messages {
		// 发送消息给用户
		err := c.Socket.WriteMessage(1, []byte(msgStr))
		if err != nil {
			return err
		}
		// 删除已发送的离线消息
		if _, err := global.RDB.LRem(global.Ctx, "offline_messages:"+c.UserID, 0, msgStr).Result(); err != nil {
			logrus.Errorf("Failed to remove offline message from Redis for %s: %v", c.UserID, err)
			// 如果删除失败，可以选择继续删除或者忽略该消息
			// 这里选择忽略该消息，下次再发送
			continue
		}
	}
	return nil
}

// func (c *Client) Write() { // 将消息发送给指定用户
//		defer func() {
//			_ = c.Socket.Close()
//		}()
//		for {
//			select {
//			// 上述start方法中，第三个case会将消息放进对应的client中，这里监听client.SendMsg通道并将消息显示出来
//			case message, ok := <-c.ReceiveMsg:
//				if !ok {
//					_ = c.Socket.WriteMessage(websocket.CloseMessage, []byte{})
//					return
//				}
//				msg, _ := json.Marshal(string(message))
//				_ = c.Socket.WriteMessage(1, msg)
//			}
//		}
//	}

// PublishRedis 向 Redis 发布消息
func (c *Client) PublishRedis(ctx context.Context, channel string, message any) {
	_, _ = global.RDB.Expire(global.Ctx, c.UserID, time.Hour*24*30*3).Result() // 建立连接3个月过期
	msg, _ := json.Marshal(message)
	global.RDB.Publish(ctx, channel, msg)
}

// SubscribeRedis 从 Redis 订阅消息，如果有新消息，则将其发送给客户端
func (c *Client) SubscribeRedis(ctx context.Context, channel string) {
	pubSub := global.RDB.Subscribe(ctx, channel)
	defer pubSub.Close()
	for {
		message, err := pubSub.ReceiveMessage(ctx)
		if err != nil {
			fmt.Println(err)
		}
		// 检查接收方是否在线
		receiverClient, ok := Manager.Clients[channel]
		if !ok || !receiverClient.Online {
			// 接收方不在线，不做处理
			continue
		}
		_ = c.Socket.WriteMessage(1, []byte(message.Payload))
	}
}
