package hcaptcha

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"time"

	"github.com/0xF7A4C6/GoCycle"
	"github.com/valyala/fasthttp"
	"github.com/zenthangplus/goccm"
)

var (
	NORMAL_ADDR = "http://127.0.0.1:4321"
)

var (
	TASKTYPE_ENTERPRISE = 0
	TASKTYPE_NORMAL     = 1
)

var (
	cc = goccm.New(1500)

	readTimeout, _  = time.ParseDuration("10s")
	writeTimeout, _ = time.ParseDuration("10s")

	headerContentTypeJson = []byte("application/json")
	Client                = &fasthttp.Client{
		ReadTimeout:                   readTimeout,
		WriteTimeout:                  writeTimeout,
		MaxIdleConnDuration:           time.Second * 15,
		NoDefaultUserAgentHeader:      true,
		DisableHeaderNamesNormalizing: true,
		DisablePathNormalizing:        true,
		Dial: (&fasthttp.TCPDialer{
			Concurrency:      4096,
			DNSCacheDuration: time.Hour,
		}).Dial,
	}
)

var (
	Endpoints = GoCycle.New(&[]string{
		"http://127.0.0.1:1235",
		"http://127.0.0.1:1236",
		"http://127.0.0.1:1237",
		"http://127.0.0.1:1238",
		"http://127.0.0.1:1239",
		"http://127.0.0.1:1240",
		"http://127.0.0.1:1241",
		"http://127.0.0.1:1242",
		"http://127.0.0.1:1243",
		"http://127.0.0.1:1244",
		"http://127.0.0.1:1245",
		"http://127.0.0.1:1246",
	})
)

func (c *Hcap) GetHsw(jwt string, isSubmit bool) (string, error) {
	for i := 0; i < 10; i++ {
		req := fasthttp.AcquireRequest()

		switch c.Config.TaskType {
		case TASKTYPE_ENTERPRISE:
			n, err := c.Manager.Build(jwt, isSubmit)
			if err != nil {
				return "", fmt.Errorf("someone poop in the api and we got a error")
			}

			out, err := json.Marshal(n)
			if err != nil {
				return "", fmt.Errorf("someone poop in the api and we got a error")
			}

			fp := base64.StdEncoding.EncodeToString(out)

			end, _ := Endpoints.Next()
			req.Header.SetMethod(fasthttp.MethodPost)
			req.Header.SetContentTypeBytes(headerContentTypeJson)
			req.SetRequestURI(fmt.Sprintf("%s/n", end))
			req.SetBodyRaw([]byte(fmt.Sprintf(`{"jwt": "%s", "fp": "%s"}`, jwt, fp)))
		case TASKTYPE_NORMAL:
			req.Header.SetMethod(fasthttp.MethodGet)
			req.SetRequestURI(fmt.Sprintf("%s/n?req=%s", NORMAL_ADDR, jwt))
		}

		resp := fasthttp.AcquireResponse()

		cc.Wait()
		err := Client.Do(req, resp)
		cc.Done()

		fasthttp.ReleaseRequest(req)
		defer fasthttp.ReleaseResponse(resp)

		if err != nil {
			continue
		}

		return string(resp.Body()), nil
	}

	return "", fmt.Errorf("someone poop in the api and we got a error")
}
