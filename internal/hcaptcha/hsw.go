package hcaptcha

import (
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

var (
	ADDR = "http://127.0.0.1:1234"
)

func (c *Hcap) GetHsw(jwt string) (string, error) {
	client := http.Client{
		Timeout: 15 * time.Second,
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/n", ADDR), strings.NewReader(fmt.Sprintf(`{"jwt": "%s"}`, jwt)))
	if err != nil {
		return "", err
	}

	req.Header.Set("content-type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil

}
