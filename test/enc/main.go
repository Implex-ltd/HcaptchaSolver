package main

import (
	"encoding/base64"
	"fmt"
	"log"
	"math/rand"
	"regexp"
	"strings"
	"time"
	"unicode/utf8"
)

var (
	charset = "abcdefghijklmnopqrstuvwxyz"
)

func Reverse(s string) string {
	size := len(s)
	buf := make([]byte, size)
	for start := 0; start < size; {
		r, n := utf8.DecodeRuneInString(s[start:])
		start += n
		utf8.EncodeRune(buf[size-start:], r)
	}
	return string(buf)
}

func findCharIndex(char byte, charset string) int {
	for i := 0; i < len(charset); i++ {
		if charset[i] == char {
			return i
		}
	}
	return -1
}

func getRandNum(A, g int) int {
	return rand.Intn(g-A+1) + A
}

func getRandomChar(min, max int) byte {
	return byte(rand.Intn(max-min+1) + min)
}

func enc(input string) []string {
	inputArr := func() string {
		var out []byte
		for i := 0; i < 13; i++ {
			char := getRandomChar(65, 90)
			out = append(out, char)
		}

		return string(out)
	}()

	randA := getRandNum(1, 26)

	Reversed := func(input string) (out string) {
		words := []string{}

		for _, word := range strings.Split(input, " ") {
			words = append(words, Reverse(word))
		}

		return strings.Join(words, " ")
	}(input)

	b64 := Reverse(base64.RawURLEncoding.WithPadding(base64.StdPadding).EncodeToString([]byte(Reversed)))
	b64rand := getRandNum(1, len(b64)-1)

	final := func(b64 string, i int) string {
		return regexp.MustCompile(fmt.Sprintf("[%v%s]", inputArr, strings.ToLower(inputArr))).ReplaceAllStringFunc(b64[i:]+b64[:i], func(A string) string {
			if A == strings.ToUpper(A) {
				return strings.ToLower(A)
			} else {
				return strings.ToUpper(A)
			}
		})
	}(b64, b64rand)

	return []string{
		final,
		fmt.Sprintf("%x", randA),
		fmt.Sprintf("%x", b64rand),
		inputArr,
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())

	for _, str := range []string{
		"Europe/Paris",
		"Google Inc. (NVIDIA)",
		"ANGLE (NVIDIA, NVIDIA GeForce RTX 3060 Ti Direct3D11 vs_5_0 ps_5_0, D3D11)",
		"143254600089",
	} {
		log.Printf("======[  %s  ]======", str)
		log.Println(enc(str))
	}
}
