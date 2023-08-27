package hcaptcha

import (
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

const (
	TYPE_ENTERPRISE = 0
	TYPE_NORMAL     = 1
	TYPE_TURBO      = 2
)

var (
	ENTERPRISE_ADDR = "http://127.0.0.1:1234"
	NORMAL_ADDR     = "http://127.0.0.1:4321"
)

func (c *Hcap) GetHsw(jwt string) (string, error) {
	client := http.Client{
		Timeout: 10 * time.Second,
	}

	var req *http.Request
	var err error

	switch c.Config.TaskType {
	case TYPE_ENTERPRISE:
		req, err = http.NewRequest("POST", fmt.Sprintf("%s/n", ENTERPRISE_ADDR), strings.NewReader(fmt.Sprintf(`{"jwt": "%s"}`, jwt)))
		if err != nil {
			return "", err
		}

		req.Header.Set("content-type", "application/json")
	case TYPE_NORMAL:
		req, err = http.NewRequest("GET", fmt.Sprintf("%s/n?req=%s", NORMAL_ADDR, jwt), nil)
		if err != nil {
			return "", err
		}
	}

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
