package utils

import (
	"crypto/md5"
	"fmt"
)

// Md5 将字符串 md5 加密
func Md5(str string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(str)))
}
