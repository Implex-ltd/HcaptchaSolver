# HcaptchaSolver
Hcaptcha enterprise challenge solver

```
	- Enterprise HSW
		port: 1234
		endpoint: /n
		body: {"jwt": "req.."}
		
	- Normal HSW
		port: 4321
		endpoint: /n?req=jwt	
```

```
docker run --rm --pull always -p 8000:8000 -v /mydata:/mydata surrealdb/surrealdb:1.0.0-beta.9-20230402 start --log trace --user root --pass root
```