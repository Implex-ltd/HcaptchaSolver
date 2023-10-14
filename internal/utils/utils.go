package utils

import (
	"fmt"
	"math/rand"
	"os"
)

const (
	charset      = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	hash_charset = "0123456789"

	precision = 10000000000000.0
	basePath  = "../../assets"
)

func RandomNumber(a, b int) int {
	if a >= b {
		panic("Invalid range: a must be less than b")
	}

	return rand.Intn(b-a+1) + a
}

func RandomFloat64Precission(a, b float64, prec float64) float64 {
	if a >= b {
		panic("Invalid range: a must be less than b")
	}

	randomFloat := a + rand.Float64()*(b-a)
	randomFloat *= prec
	randomFloat = float64(int(randomFloat)) / prec

	return randomFloat
}

func RandomFloat64(a, b float64) float64 {
	if a >= b {
		panic("Invalid range: a must be less than b")
	}

	randomFloat := a + rand.Float64()*(b-a)
	randomFloat *= precision
	randomFloat = float64(int(randomFloat)) / precision

	return randomFloat
}

func RandomString(length int) string {
	result := make([]byte, length)
	for i := range result {
		result[i] = charset[rand.Intn(len(charset))]
	}

	return string(result)
}

func RandomHash(length int) string {
	result := make([]byte, length)
	for i := range result {
		result[i] = hash_charset[rand.Intn(len(hash_charset))]
	}

	return string(result)
}

func RandomElementInt(slice []int) int {
	index := rand.Intn(len(slice))
	return slice[index]
}

func RandomElementString(slice []string) string {
	index := rand.Intn(len(slice))
	return slice[index]
}

func AppendLine(text, fileName string) error {
	file, err := os.OpenFile(fmt.Sprintf("%s/%s", basePath, fileName), os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	if _, err := file.WriteString(text + "\n"); err != nil {
		return err
	}

	return nil
}
