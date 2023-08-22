import httpx
from rich import print

res = httpx.post(
    "http://127.0.0.1:1337/solve",
    json={
        "domain": "discord.com",
        "site_key": "4c672d35-0701-42b2-88c3-78380b0db560",
        "user_agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/116.0.0.0 Safari/537.36",
        "proxy": "http://user:pass@ip:port",
    },
    timeout=None,
)

print(res.status_code, res.json())
