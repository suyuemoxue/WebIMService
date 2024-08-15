package receive

// RegisterReceiveStruct 用于接收注册信息
type RegisterReceiveStruct struct {
	UserName   string `json:"username" form:"username" binding:"required"`
	Password   string `json:"password" form:"password" binding:"required"`
	RePassword string `json:"rePassword" form:"rePassword" binding:"required"`
	//VerificationCode string `json:"verificationCode" form:"verificationCode" binding:"required"`
}

// LoginReceiveStruct 用于接收登录信息
type LoginReceiveStruct struct {
	UserName string `json:"username" form:"username" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
}

type GetHistoryMsg struct {
	UserName   string `json:"username" form:"username" binding:"required"`
	TargetName string `json:"targetname" form:"targetname" binding:"required"`
}
