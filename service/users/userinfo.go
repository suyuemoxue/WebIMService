package users

import (
	"WebIM/global"
	"WebIM/global/receive"
	"WebIM/global/send"
	"WebIM/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserInfoService struct {
}

// GetUserList
// @Tags 查询所有用户
// @Success 200 {string} json{"code","message"}
// @Router /getUserList [get]
func (us *UserInfoService) GetUserList(context *gin.Context) {
	user := models.User{}
	userList := user.GetUserList()
	context.JSON(http.StatusOK, global.Response{
		Code: http.StatusOK,
		Data: userList,
		Msg:  "查询成功",
	})
}

// GetUserInfo
// @Tags 根据条件查询用户信息
// @Success 200 {string} json{"code","message"}
// @Router /getUserInfo [get]
// GetUserInfo 获取用户信息
func (us *UserInfoService) GetUserInfo(context *gin.Context) {
	user := models.User{}
	if !user.CheckUserExists("name", context.PostForm("name")) {
		context.JSON(-1, global.Response{
			Code: -1,
			Data: nil,
			Msg:  "用户名不存在",
		})
		return
	}
	if !user.GetUserInfo("name", context.PostForm("name")) {
		context.JSON(-1, global.Response{
			Code: -1,
			Data: nil,
			Msg:  "查询失败",
		})
		return
	}
	context.JSON(http.StatusOK, global.Response{
		Code: http.StatusOK,
		Data: user,
		Msg:  "查询成功",
	})
}

// GetUserHistoryMessage
// @Tags 查询历史消息
// @Success 200 {string} json{"code","message"}
// @Router /getUserHistoryMessage [get]
// GetUserInfo 获取用户信息
func (us *UserInfoService) GetUserHistoryMessage(context *gin.Context) {
	msg := models.Message{}
	history := receive.GetHistoryMsg{}
	err := context.ShouldBind(&history)
	if err != nil {
		return
	}
	historyMsg, num := msg.GetHistoryMsg(history.UserName, history.TargetName)
	for _, message := range historyMsg {
		returnMsg := send.ReturnMsg{
			SendID:    message.SendID,
			ReceiveID: message.ReceiveID,
			Content:   string(message.Content),
			MsgType:   message.MsgType,
			MediaType: message.MediaType,
		}
		if num > 1 {
			context.JSON(http.StatusOK, returnMsg)     // 发送当前 JSON 响应的分块
			_, err := context.Writer.WriteString("\n") // 添加分隔符
			if err != nil {
				return
			}
			// 强制刷新响应缓冲区，将数据发送给客户端
			context.Writer.Flush()
		} else {
			context.JSON(http.StatusOK, returnMsg)
		}
	}
}
