package users

import (
	"WebIM/service/users"
	"github.com/gin-gonic/gin"
)

type UserRouter struct {
}

func (ur *UserRouter) InitUserRouter(router *gin.RouterGroup) {
	userRouter := router.Group("")
	{
		userService := users.UserInfoService{}
		userRouter.POST("/getUserList", userService.GetUserList)
		userRouter.POST("/getUserInfo", userService.GetUserInfo)
		userRouter.POST("/getUserHistoryMessage", userService.GetUserHistoryMessage)
	}
}
