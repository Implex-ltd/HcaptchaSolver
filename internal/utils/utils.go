package utils

import "math/rand"

const (
	charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
)

func RandomNumber(a, b int) int {
	if a >= b {
		panic("Invalid range: a must be less than b")
	}

	return rand.Intn(b-a+1) + a
}

func RandomString(length int) string {
	result := make([]byte, length)
	for i := range result {
		result[i] = charset[rand.Intn(len(charset))]
	}

	return string(result)
}