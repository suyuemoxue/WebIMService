package users

import (
	"WebIM/service/users"
	"github.com/gin-gonic/gin"
)

type LoginRouter struct {
}

func (lr *LoginRouter) InitLoginRouter(router *gin.RouterGroup) {
	loginRouter := router.Group("")
	{
		loginService := users.LoginService{}
		loginRouter.POST("/login", loginService.Login)
		loginRouter.POST("/register", loginService.Register)
	}
}
