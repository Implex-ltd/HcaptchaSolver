package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"sync"
	"time"
)

const (
	STATUS_SOLVING = 0
	STATUS_SOLVED  = 1
	STATUS_ERROR   = 2
)

type BodyNewSolveTask struct {
	Domain    string `json:"domain"`
	SiteKey   string `json:"site_key"`
	UserAgent string `json:"user_agent"`
	Proxy     string `json:"proxy"`
}

type TaskResponse struct {
	Data struct {
		CreatedAt  time.Time `json:"CreatedAt"`
		UpdatedAt  time.Time `json:"UpdatedAt"`
		DeletedAt  time.Time `json:"DeletedAt"`
		ID         string    `json:"ID"`
		Status     int       `json:"status"`
		Token      string    `json:"token"`
		Error      string    `json:"error"`
		Success    bool      `json:"success"`
		Expiration int       `json:"expiration"`
	} `json:"data"`
	Message string `json:"message"`
	Success bool   `json:"success"`
}

type CheckResponse struct {
	Data struct {
		CreatedAt  time.Time `json:"CreatedAt"`
		UpdatedAt  time.Time `json:"UpdatedAt"`
		DeletedAt  any       `json:"DeletedAt"`
		ID         string    `json:"ID"`
		Status     int       `json:"status"`
		Token      string    `json:"token"`
		Error      string    `json:"error"`
		Success    bool      `json:"success"`
		Expiration int       `json:"expiration"`
	} `json:"data"`
	Message string `json:"message"`
	Status  string `json:"status"`
}

func solve(proxy string) (string, error) {
	payload, _ := json.Marshal(BodyNewSolveTask{
		UserAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/113.0.0.0 Safari/537.36",
		Domain:    "accounts.hcaptcha.com",
		SiteKey:   "2eaf963b-eeab-4516-9599-9daa18cd5138",
		Proxy:     proxy,
	})

	// create task
	client := http.Client{
		Timeout: 15 * time.Second,
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/api/task/new", "http://127.0.0.1:3000"), strings.NewReader(string(payload)))
	if err != nil {
		return "", err
	}
	req.Header.Add("content-type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	fmt.Println(string(body))

	var out TaskResponse
	if err := json.Unmarshal(body, &out); err != nil {
		return "", err
	}

	if !out.Success {
		return "", fmt.Errorf(out.Message)
	}

	log.Printf("Created task (%v)\n", out.Data.ID)

	// get result
	for {
		req, err := http.NewRequest("GET", fmt.Sprintf("%s/api/task/%s", "http://127.0.0.1:3000", out.Data.ID), strings.NewReader(string(payload)))
		if err != nil {
			return "", nil
		}
		req.Header.Add("content-type", "application/json")

		resp, err := client.Do(req)
		if err != nil {
			return "", err
		}

		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return "", err
		}

		var out CheckResponse
		if err := json.Unmarshal(body, &out); err != nil {
			return "", err
		}

		switch out.Data.Status {
		case STATUS_ERROR:
			return "", fmt.Errorf(out.Message)
		case STATUS_SOLVED:
			return out.Data.Token, nil
		case STATUS_SOLVING:
			time.Sleep(1 * time.Second)
		}
	}
}

func main() {
	w := sync.WaitGroup{}
	w.Add(5)

	for i := 0; i < 5; i++ {
		go func() {
			defer w.Done()

			token, err := solve("http://brd-customer-hl_5ae0707e-zone-data_center-ip-178.171.117.118:s3a3gvzzhgt8@brd.superproxy.io:22225")
			if err != nil {
				log.Println(err)
				return
			}

			fmt.Printf("task solved: %s", token[:25])
		}()
	}

	w.Wait()
}
