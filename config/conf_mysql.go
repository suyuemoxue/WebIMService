package config

import (
	"strings"
)

type MySQL struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	DB       string `yaml:"db"`
	User     string `yaml:"users"`
	Password string `yaml:"password"`
	LogLevel string `yaml:"log_level"` //日志等级
}

func (ms *MySQL) Dsn() (dsn string) {
	var sb strings.Builder
	// 高效拼接
	sb.WriteString(ms.User + ":" + ms.Password + "@tcp(" + ms.Host + ":" + ms.Port + ")/" + ms.DB + "?charset=utf8&parseTime=True&loc=Local")
	dsn = sb.String()
	return
}
