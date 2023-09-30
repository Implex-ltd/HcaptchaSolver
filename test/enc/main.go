package main

import (
	"encoding/base64"
	"fmt"
	"math/rand"
	"net/url"
	"strings"
	"time"
	"unicode"
)

var (
	charset = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
)

func swapCase(input string) string {
	return strings.Map(func(r rune) rune {
		if r >= 'A' && r <= 'Z' {
			return r + 32
		} else if r >= 'a' && r <= 'z' {
			return r - 32
		}
		return r
	}, input)
}

func reverseString(input string) string {
	runes := []rune(input)
	reversed := make([]rune, len(runes))
	for i, j := 0, len(runes)-1; i <= j; i, j = i+1, j-1 {
		reversed[i], reversed[j] = runes[j], runes[i]
	}
	return string(reversed)
}

func charCodeAt(s string, n int) rune {
	i := 0
	for _, r := range s {
		if i == n {
			return r
		}
		i++
	}
	return 0
}

func uppercaseRune(char rune) rune {
    if !unicode.IsLower(char) {
        return char
    }
    return unicode.ToUpper(char)
}

func encryptStr(input string) ([]string, error) {
	if input == "" {
		return nil, fmt.Errorf("Input cannot be empty")
	}

	randomOffset := rand.Intn(26)

	encryptedString := func(input string) string {
		words := strings.Split(input, " ")
		encWords := make([]string, len(words))

		for _, word := range words {
			reversed := reverseString(word)
			enc := ""

			for _, char := range reversed {
				if !strings.ContainsRune(charset, uppercaseRune(char)) {
					enc += string(char)
					continue
				}

				charCode := int(charCodeAt(strings.ToLower(fmt.Sprintf("%c", char)), 0))
				encryptedCharCode := (charCode-97+randomOffset)%26 + 97

				if char == uppercaseRune(char) {
					enc += fmt.Sprintf("%c", encryptedCharCode-32)
				} else {
					enc += fmt.Sprintf("%c", encryptedCharCode)
				}
			}

			encWords = append(encWords, enc)
		}

		return strings.Join(encWords, " ")[3:]
	}(input)

	encodedString := func(encrypted string) string {
		b64 := base64.RawStdEncoding.EncodeToString([]byte(url.QueryEscape(encrypted)))
		return swapCase(b64)
	}(encryptedString)

	length := len(encodedString)
	randomIndex := rand.Intn(length-1) + 1

	firstPart := encodedString[randomIndex:]
	secondPart := encodedString[:randomIndex]

	finalString := (firstPart + secondPart)
	finalString = swapCase(finalString)

	return []string{finalString, fmt.Sprintf("%x", randomOffset), fmt.Sprintf("%x", randomIndex)}, nil
}

func main() {
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < 5; i++ {
		result, err := encryptStr("Europe/Paris")
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		fmt.Println(result)
	}
}
