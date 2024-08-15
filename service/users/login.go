package users

import (
	"WebIM/global"
	"WebIM/global/receive"
	"WebIM/models"
	"WebIM/utils"
	"fmt"
	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"math/rand"
	"net/http"
	"time"
)

type LoginService struct {
}

// Login
// @Tags 登录
// @Success 200 {string} json{"code","message"}
// @Router /login [post]
func (ls *LoginService) Login(context *gin.Context) {
	user := models.User{}
	login := receive.LoginReceiveStruct{}
	err := context.ShouldBind(&login)
	if err != nil {
		return
	}
	// 判断用户名是否存在
	if !user.CheckUserExists("name", login.UserName) {
		context.JSON(http.StatusOK, global.Response{
			Code: -1,
			Data: nil,
			Msg:  "该用户名不存在",
		})
		return
	}
	fmt.Println(login)
	// 判断密码是否正确
	if !user.CheckPasswordByUsername(login.UserName, login.Password) {
		context.JSON(http.StatusOK, global.Response{
			Code: -1,
			Data: nil,
			Msg:  "密码错误，请重新输入",
		})
		return
	}
	context.JSON(http.StatusOK, global.Response{
		Code: http.StatusOK,
		Data: user,
		Msg:  "登录成功",
	})
}

// Register
// @Tags 注册
// @Success 200 {string} json{"code","message"}
// @Router /register [post]
func (ls *LoginService) Register(context *gin.Context) {
	user := models.User{}
	register := receive.RegisterReceiveStruct{}
	// 检查是否缺少数据
	if err := context.ShouldBind(&register); err != nil {
		context.JSON(http.StatusOK, global.Response{
			Code: -1,
			Data: register,
			Msg:  "请检查是否缺少数据",
		})
		return
	}
	// 检查两次密码是否一致
	if register.Password != register.RePassword {
		context.JSON(http.StatusOK, global.Response{
			Code: -1,
			Data: nil,
			Msg:  "两次密码不一致，请重新输入",
		})
		return
	}
	// 检查用户名是否已存在
	if user.CheckUserExists("name", register.UserName) {
		context.JSON(http.StatusOK, global.Response{
			Code: -1,
			Data: nil,
			Msg:  "该用户名已存在，无需注册",
		})
		return
	}
	user.Name = register.UserName
	// 密码加密
	salt := make([]byte, 6)
	for i := range salt {
		salt[i] = utils.SaltStr[rand.Int63()%int64(len(utils.SaltStr))]
	}
	user.Password = utils.MakePassword(register.Password, string(salt))
	user.Salt = string(salt)
	user.LoginTime = time.Now()
	user.LoginOutTime = time.Now()
	user.HeartbeatTime = time.Now()
	// 添加数据
	_, err := govalidator.ValidateStruct(user)
	if err != nil {
		return
	}
	if !user.CreateUser() {
		context.JSON(http.StatusOK, global.Response{
			Code: -1,
			Data: nil,
			Msg:  "注册失败",
		})
		return
	}
	context.JSON(http.StatusOK, global.Response{
		Code: http.StatusOK,
		Data: user,
		Msg:  "注册成功",
	})
}
