package utils

import (
	"math/rand"
	"time"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func InitRandom() {
	rand.Seed(time.Now().UnixNano())
}

func RandSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func RandBool() bool {
	return rand.Intn(2) == 1
}

func RandInt(min int, max int) int {
	return rand.Intn(max-min) + min
}
