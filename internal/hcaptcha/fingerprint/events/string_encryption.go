package events

import (
	"encoding/base64"
	"fmt"
	"html"
	"math/rand"
	"net/url"
	"regexp"
	"strconv"
	"strings"
	"unicode/utf8"
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

func getRandNum(A, g int) int {
	return rand.Intn(g-A+1) + A
}

func getRandomChar(min, max int) byte {
	return byte(rand.Intn(max-min+1) + min)
}

func EncStr(input string) []string {
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

	b64 := Reverse(base64.StdEncoding.EncodeToString([]byte(url.QueryEscape(Reversed))))
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

func decode(s string) string {
    if strings.IndexByte(s, ';') >= 0 {
        s = html.UnescapeString(s)
    }
    return s
}

func DecStr(input []string) string {
	b64rand, _ := strconv.ParseInt(input[2], 16, 64)
	inputArr := input[3]
	final := input[0]

	// Step 1: Reverse character substitution
	decrypted := regexp.MustCompile(fmt.Sprintf("[%v%s]", inputArr, strings.ToLower(inputArr))).ReplaceAllStringFunc(final, func(A string) string {
		if A == strings.ToUpper(A) {
			return strings.ToLower(A)
		} else {
			return strings.ToUpper(A)
		}
	})

	// Step 2: Reverse the random Base64 offset
	b64Len := len(decrypted)
	reversedB64 := decrypted[b64Len-int(b64rand):] + decrypted[:b64Len-int(b64rand)]
	decoded, _ := base64.StdEncoding.DecodeString(Reverse(reversedB64))

	// Step 3: Reverse word reversal
	reversedStr := decode(Reverse(string(decoded)))
	return reversedStr
}