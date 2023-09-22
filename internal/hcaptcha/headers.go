package hcaptcha

import http "github.com/bogdanfinn/fhttp"

var (
	order = []string{
		`authority`,
		`accept`,
		`accept-language`,
		`cache-control`,
		`content-length`,
		`content-type`,
		`origin`,
		`pragma`,
		`referer`,
		`sec-ch-ua`,
		`sec-ch-ua-mobile`,
		`sec-ch-ua-platform`,
		`sec-fetch-dest`,
		`sec-fetch-mode`,
		`sec-fetch-site`,
		`user-agent`,
	}
)

func (c *Hcap) HeaderCheckSiteConfig() http.Header {
	return http.Header{
		`authority`:          {`hcaptcha.com`},
		`accept`:             {`application/json`},
		`accept-language`:    {c.Http.BaseHeader.AcceptLanguage},
		`cache-control`:      {`no-cache`},
		`content-type`:       {`text/plain`},
		`origin`:             {`https://newassets.hcaptcha.com`},
		`pragma`:             {`no-cache`},
		`referer`:            {`https://newassets.hcaptcha.com/`},
		`sec-ch-ua`:          {c.Http.BaseHeader.SecChUa},
		`sec-ch-ua-mobile`:   {c.Http.BaseHeader.SecChUaMobile},
		`sec-ch-ua-platform`: {c.Http.BaseHeader.SecChUaPlatform},
		`sec-fetch-dest`:     {`empty`},
		`sec-fetch-mode`:     {`cors`},
		`sec-fetch-site`:     {`same-site`},
		`user-agent`:         {c.Fingerprint.Navigator.UserAgent},

		http.HeaderOrderKey: order,
	}
}

func (c *Hcap) HeaderGetCaptcha() http.Header {
	return http.Header{
		`authority`:          {`hcaptcha.com`},
		`accept`:             {`application/json`},
		`accept-language`:    {c.Http.BaseHeader.AcceptLanguage},
		`cache-control`:      {`no-cache`},
		`content-type`:       {`application/x-www-form-urlencoded`},
		`origin`:             {`https://newassets.hcaptcha.com`},
		`pragma`:             {`no-cache`},
		`referer`:            {`https://newassets.hcaptcha.com/`},
		`sec-ch-ua`:          {c.Http.BaseHeader.SecChUa},
		`sec-ch-ua-mobile`:   {c.Http.BaseHeader.SecChUaMobile},
		`sec-ch-ua-platform`: {c.Http.BaseHeader.SecChUaPlatform},
		`sec-fetch-dest`:     {`empty`},
		`sec-fetch-mode`:     {`cors`},
		`sec-fetch-site`:     {`same-site`},
		`user-agent`:         {c.Fingerprint.Navigator.UserAgent},

		http.HeaderOrderKey: order,
	}
}

func (c *Hcap) HeaderCheckCaptcha() http.Header {
	return http.Header{
		`authority`:          {`hcaptcha.com`},
		`accept`:             {`*/*`},
		`accept-language`:    {c.Http.BaseHeader.AcceptLanguage},
		`cache-control`:      {`no-cache`},
		`content-type`:       {`application/json;charset=UTF-8`},
		`origin`:             {`https://newassets.hcaptcha.com`},
		`pragma`:             {`no-cache`},
		`referer`:            {`https://newassets.hcaptcha.com/`},
		`sec-ch-ua`:          {c.Http.BaseHeader.SecChUa},
		`sec-ch-ua-mobile`:   {c.Http.BaseHeader.SecChUaMobile},
		`sec-ch-ua-platform`: {c.Http.BaseHeader.SecChUaPlatform},
		`sec-fetch-dest`:     {`empty`},
		`sec-fetch-mode`:     {`cors`},
		`sec-fetch-site`:     {`same-site`},
		`user-agent`:         {c.Fingerprint.Navigator.UserAgent},

		http.HeaderOrderKey: order,
	}
}
