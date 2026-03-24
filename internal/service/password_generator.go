package service

import (
	"crypto/rand"
	"math/big"
)

const length = 10

var possibleRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&-)(*")

func GenerateRandomPassword() (string, error) {
	password := make([]rune, length)

	for i := 0; i < length; i++ {
		n, err := rand.Int(rand.Reader, big.NewInt(int64(len(possibleRunes))))
		if err != nil {
			return "", err
		}
		
		password[i] = possibleRunes[n.Int64()]
	}

	return string(password), nil
}
