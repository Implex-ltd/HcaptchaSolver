package hcaptcha

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

var (
	ADDR = "http://127.0.0.1:4000"
)

func (c *Hcap) GetHsw(jwt string) (string, error) {
	// req, err := http.NewRequest("POST", "https://lkl9pn-ip-88-180-215-157.tunnelmole.net/n", strings.NewReader(fmt.Sprintf(`{"jwt": "%s"}`, jwt)))

	for {
	client := http.Client{
		Timeout: 10 * time.Second,
	}

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/n?req=%s", ADDR, jwt), nil)
	if err != nil {
		continue//return "", err
	}

	//req.Header.Set("content-type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err.Error())
		continue//return "", err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		continue//return "", err
	}

	return string(body), nil
}
}
