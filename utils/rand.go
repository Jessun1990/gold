package utils

import (
	"math/rand"
	"time"
)

func GetTrueRandIntn(num int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(num)
}
