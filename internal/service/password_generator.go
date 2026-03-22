package service

import (
	"math/rand"
	"time"
)

const length = 8

var possibleRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&-)(*")
var countPossibleRunes int32 = int32(len(possibleRunes))

func generateRandomPassword() string {
	password := make([]rune, length)

	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	for i := 0; i < length; i++ {
		password[i] = possibleRunes[r.Int31n(countPossibleRunes)]
	}

	return string(password)
}
