package hcaptcha

import (
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/zenthangplus/goccm"
)

var (
	ENTERPRISE_ADDR = "http://127.0.0.1:1234"
	NORMAL_ADDR     = "http://127.0.0.1:4321"
	cc              = goccm.New(150) // 200
)

var (
	TASKTYPE_ENTERPRISE = 0
	TASKTYPE_NORMAL     = 1
)

var Client *http.Client

func (c *Hcap) GetHsw(jwt string) (string, error) {
	var req *http.Request
	var err error

	switch c.Config.TaskType {
	case TASKTYPE_ENTERPRISE:
		req, err = http.NewRequest("POST", fmt.Sprintf("%s/n", ENTERPRISE_ADDR), strings.NewReader(fmt.Sprintf(`{"jwt": "%s"}`, jwt)))
		if err != nil {
			return "", err
		}

		req.Header.Set("content-type", "application/json")
	case TASKTYPE_NORMAL:
		req, err = http.NewRequest("GET", fmt.Sprintf("%s/n?req=%s", NORMAL_ADDR, jwt), nil)
		if err != nil {
			return "", err
		}
	}

	cc.Wait()
	resp, err := Client.Do(req)
	cc.Done()
	if err != nil {
		return "", err
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println("hsw", err)
		}
	}(resp.Body)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}
