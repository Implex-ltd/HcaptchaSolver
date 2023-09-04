import httpx, time, threading

__server__ = "127.0.0.1:3000"


class STATUS:
    STATUS_SOLVING = 0
    STATUS_SOLVED = 1
    STATUS_ERROR = 2


class TASK_TYPE:
    TYPE_ENTERPRISE = 0
    """
    enterprise anti fingerprinting browser engine.
    """

    TYPE_NORMAL = 1
    """
    jsdom hsw sandboxing
    """


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
        useragent: str = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/115.0.0.0 Safari/537.36",
        proxy: str = "",
        invisible: bool = False,
        rqdata: str = "",
        text_free_entry: bool = False,
        turbo: bool = False,
        turbo_st: int = 3000,
    ) -> dict:
        """
        Create a new task for solving a captcha.

        Args:
            `task_type` (TASK_TYPE, optional): The type of captcha-solving task. Defaults to TASK_TYPE.TYPE_NORMAL.

            `domain` (str, optional): The domain where the captcha is presented. Defaults to "accounts.hcaptcha.com".

            `sitekey` (str, optional): The sitekey associated with the captcha. Defaults to "2eaf963b-eeab-4516-9599-9daa18cd5138".

            `useragent` (str, optional): The user agent to use when making requests. Defaults to a common user agent string.

            `proxy` (str, optional): The proxy to use for making requests. Defaults to an empty string.

            `invisible` (bool, optional): Whether the captcha is invisible. Defaults to False.

            `rqdata` (str, optional): Additional request data. Defaults to an empty string.

            `text_free_entry` (bool, optional): Whether free text entry is allowed. Defaults to False.

            `turbo` (bool, optional): Whether turbo mode is enabled. Defaults to False.

            `turbo_st` (int, optional): The turbo mode submit time in milliseconds. Defaults to 3000 (3s).

        Returns:
            dict: A dictionary containing the task ID.
        """
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
                    "a11y_tfe": text_free_entry,
                    "turbo": turbo,
                    "turbo_st": turbo_st,
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
        domain="discord.com",
        sitekey="4c672d35-0701-42b2-88c3-78380b0db560",
        turbo_st=100,
        turbo=True,
        proxy="http://brd-customer-hl_5ae0707e-zone-data_center-ip-45.134.115.190:s3a3gvzzhgt8@brd.superproxy.io:22225",
        text_free_entry=True,
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
    for _ in range(10):
        threading.Thread(target=task).start()
