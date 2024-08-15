package config

import (
	"strconv"
	"strings"
)

type System struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
	Env  string `yaml:"env"`
}

func (s *System) Address() (address string) {
	var sb strings.Builder
	sb.WriteString(s.Host + ":" + strconv.Itoa(s.Port))
	address = sb.String()
	return
}
