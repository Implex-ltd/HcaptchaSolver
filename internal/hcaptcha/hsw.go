package hcaptcha

import (
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/zenthangplus/goccm"
)

var (
	ENTERPRISE_ADDR = "http://127.0.0.1:1234"
	NORMAL_ADDR     = "http://127.0.0.1:4321"
	cc              = goccm.New(150)
)

var Client *http.Client

func (c *Hcap) GetHsw(jwt string) (string, error) {
	var req *http.Request
	var err error

	switch c.Config.TaskType {
	case 0: // ENTERPRISE
		req, err = http.NewRequest("POST", fmt.Sprintf("%s/n", ENTERPRISE_ADDR), strings.NewReader(fmt.Sprintf(`{"jwt": "%s"}`, jwt)))
		if err != nil {
			return "", err
		}

		req.Header.Set("content-type", "application/json")
	case 1: // NORMAL
		req, err = http.NewRequest("GET", fmt.Sprintf("%s/n?req=%s", NORMAL_ADDR, jwt), nil)
		if err != nil {
			return "", err
		}
	}

	for i := 0; i < 3; i++ {
		cc.Wait()
		resp, err := Client.Do(req)
		cc.Done()
		if err != nil {
			time.Sleep(1 * time.Second)
			continue //return "", err
		}

		defer func(Body io.ReadCloser) {
			err := Body.Close()
			if err != nil {
				fmt.Println("hsw", err)
			}
		}(resp.Body)

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			continue //return "", err
		}

		b := string(body)

		if b == "Gay" {
			time.Sleep(3 * time.Second)
			return "", fmt.Errorf("cant get hsw")
		}

		return b, nil
	}

	return "", fmt.Errorf("hsw max retry reached")
}
