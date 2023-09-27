package hcaptcha

import (
	"fmt"
	"time"

	"github.com/valyala/fasthttp"
	"github.com/zenthangplus/goccm"
)

var (
	ENTERPRISE_ADDR = "http://127.0.0.1:1234"
	NORMAL_ADDR     = "http://127.0.0.1:4321"
)

var (
	TASKTYPE_ENTERPRISE = 0
	TASKTYPE_NORMAL     = 1
)

var (
	cc = goccm.New(150)

	readTimeout, _  = time.ParseDuration("15s")
	writeTimeout, _ = time.ParseDuration("15s")

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

func (c *Hcap) GetHsw(jwt string) (string, error) {
	req := fasthttp.AcquireRequest()

	switch c.Config.TaskType {
	case TASKTYPE_ENTERPRISE:
		req.Header.SetMethod(fasthttp.MethodPost)
		req.Header.SetContentTypeBytes(headerContentTypeJson)
		req.SetRequestURI(fmt.Sprintf("%s/n", ENTERPRISE_ADDR))
		req.SetBodyRaw([]byte(fmt.Sprintf(`{"jwt": "%s"}`, jwt)))
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
		return "", err
	}

	return string(resp.Body()), nil
}