package utils

import (
	"math/rand"
	"time"
)

// GetTrueRandIntn 获取随机整数
func GetTrueRandIntn(num int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(num)
}
