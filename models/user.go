package models

import (
	"WebIM/global"
	"WebIM/utils"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	Identity      string    `gorm:"size:10" json:"identity"`
	Name          string    `gorm:"size:10" json:"name"`
	Salt          string    `gorm:"salt" json:"salt"`
	Password      string    `gorm:"size:50" json:"password"`
	Phone         string    `gorm:"size:10" json:"phone" valid:"matches(^1[3-9]{1}\\d{9}$)"`
	Email         string    `gorm:"size:10" json:"email" valid:"email"`
	ClientIp      string    `gorm:"size:10" json:"clientIp"`
	ClientPort    string    `gorm:"size:10" json:"clientPort"`
	LoginTime     time.Time `gorm:"size:3" json:"loginTime"`
	LoginOutTime  time.Time `gorm:"size:3" json:"loginOutTime"`
	HeartbeatTime time.Time `gorm:"size:3" json:"heartbeatTime"`
	IsLogOut      bool      `gorm:"size:10" json:"isLogOut"`
	DeviceInfo    string    `gorm:"size:10" json:"deviceInfo"`
}

type UserList []User

func (u *User) TableName() string {
	return "user"
}

// GetUserInfo 查询用户信息
func (u *User) GetUserInfo(field string, value any) bool {
	err := global.DB.Debug().Where(field, value).Take(&u).Error
	if err != nil {
		return false
	}
	return true
}

// GetUserList 查找全部用户
func (u *User) GetUserList() (userList UserList) {
	userList = make(UserList, 10)
	global.DB.Debug().Find(&userList)
	return
}

// CreateUser 添加一个用户
func (u *User) CreateUser() bool {
	// 生成identity
	str := fmt.Sprintf("%d", time.Now().Unix())
	str = utils.MD5Encode(str)
	u.Identity = str[0:10]
	err := global.DB.Debug().Create(&u).Error
	if err != nil {
		return false
	}
	return true
}

// CheckUserExists 根据字段检查用户是否存在
func (u *User) CheckUserExists(field string, value any) bool {
	var count int64
	err := global.DB.Debug().Table("user").Where(field, value).Count(&count).Error
	if err != nil {
		fmt.Println(err)
		return false
	}
	return count > 0
}

// CheckPasswordByUsername 根据用户名判断密码是否正确
func (u *User) CheckPasswordByUsername(username string, password string) bool {
	err := global.DB.Debug().Where("name=?", username).First(&u).Error
	if err != nil {
		return false
	}
	res := utils.MakePassword(password, u.Salt)
	if u.Password != res {
		return false
	}
	return true
}

// ChangeUserInfo 更新用户信息
func (u *User) ChangeUserInfo(field any) bool {
	//tx := global.DB.Debug().Where("identity=?", u.Identity).Update("", field)
	return true
}

// 删除用户（软删除）
//
