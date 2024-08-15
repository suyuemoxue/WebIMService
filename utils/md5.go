package utils

import (
	"crypto/md5"
	"encoding/hex"
	"strings"
)

var SaltStr = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

// Md5Encode 小写
func Md5Encode(data string) string {
	h := md5.New()
	h.Write([]byte(data))
	sum := h.Sum(nil)
	return hex.EncodeToString(sum)
}

// MD5Encode 大写
func MD5Encode(data string) string {
	return strings.ToUpper(Md5Encode(data))
}

// MakePassword 加密
func MakePassword(plainPwd, salt string) string {
	return Md5Encode(plainPwd + salt)
}
