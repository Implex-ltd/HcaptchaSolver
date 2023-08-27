import httpx, time, threading
from rich import print

__server__ = "127.0.0.1:3000"


class STATUS:
    STATUS_SOLVING = 0
    STATUS_SOLVED = 1
    STATUS_ERROR = 2


class Api:
    def __init__(self, api_key: str = None, user_id: str = None):
        self.client = httpx.Client()
        self.user_id = user_id
        self.api_key = api_key

    def check_response(self, data: httpx.Response):
        return {
            "status": data.status_code,
            "json": data.json(),
        }

    def register(self):
        response = self.check_response(
            self.client.post(
                f"http://{__server__}/api/user/register",
            )
        )

        if response.get("json")["success"]:
            self.user_id = response["json"]["data"]["ID"]
            self.api_key = response["json"]["data"]["api_key"]

        return response

    def new_task(self):
        response = self.check_response(
            self.client.post(
                f"http://{__server__}/api/task/new",
                json={
                    "domain": "discord.com",
                    "site_key": "4c672d35-0701-42b2-88c3-78380b0db560",
                    "user_agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/116.0.0.0 Safari/537.36",
                    "proxy": "http://brd-customer-hl_5ae0707e-zone-data_center-ip-178.171.117.118:s3a3gvzzhgt8@brd.superproxy.io:22225",
                },
            )
        )

        return response

    def get_task(self, task_id: str):
        response = self.check_response(
            self.client.get(
                f"http://{__server__}/api/task/{task_id}",
            )
        )

        return response

def task():
    wrapper = Api()

    task = wrapper.new_task()
    print(task)

    token = ""
    while True:
        status = wrapper.get_task(task['json']['data'][0]['id'])
        data = status['json']['data']
        print(data)
        
        if data['status'] == STATUS.STATUS_ERROR:
            print('task failed')
            break

        elif data['status'] == STATUS.STATUS_SOLVED:
            token = data['token']
            break
        
        else:
            print('solving..')
            time.sleep(1)
    
    if token != "":
        print(f'[+] Solved: {token[:50]}')

if __name__ == "__main__":
    for _ in range(5):
        threading.Thread(target=task).start()