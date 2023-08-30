import httpx, time, threading

__server__ = "127.0.0.1:3000"


class STATUS:
    STATUS_SOLVING = 0
    STATUS_SOLVED = 1
    STATUS_ERROR = 2


class TASK_TYPE:
    TYPE_ENTERPRISE = 0
    TYPE_NORMAL = 1
    TYPE_TURBO = 2


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

    def new_task(
        self,
        task_type: TASK_TYPE = TASK_TYPE.TYPE_NORMAL,
        domain: str = "accounts.hcaptcha.com",
        sitekey: str = "2eaf963b-eeab-4516-9599-9daa18cd5138",
        useragent: str = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/116.0.0.0 Safari/537.36",
        proxy: str = "",
        invisible: bool = False,
        rqdata: str = "",
    ):
        response = self.check_response(
            self.client.post(
                f"http://{__server__}/api/task/new",
                json={
                    "domain": domain,
                    "site_key": sitekey,
                    "user_agent": useragent,
                    "proxy": proxy,
                    "task_type": task_type,
                    "invisible": invisible,
                    "rqdata": rqdata,
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

    task = wrapper.new_task(
        task_type=TASK_TYPE.TYPE_ENTERPRISE,
        domain="www.habbo.fr",
        sitekey="edc4ce89-8903-4906-80b1-7440ad9a69c8",
        invisible=True,
        proxy="http://u:p@ip:port"
    )
    print(task)

    token = ""
    while True:
        status = wrapper.get_task(task["json"]["data"][0]["id"])
        data = status["json"]["data"]
        print(data)

        if data["status"] == STATUS.STATUS_ERROR:
            break

        elif data["status"] == STATUS.STATUS_SOLVED:
            token = data["token"]
            break

        else:
            time.sleep(1)

    if token != "":
        print(f"[+] Solved: {token[:50]}")


if __name__ == "__main__":
    for _ in range(5):
        threading.Thread(target=task).start()
