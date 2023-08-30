import httpx, time, threading

__server__ = "127.0.0.1:3000"

"""
        # one-click
        task = wrapper.new_task(
            task_type=TASK_TYPE.TYPE_NORMAL,
            domain="www.habbo.fr",
            sitekey="edc4ce89-8903-4906-80b1-7440ad9a69c8",
        )

        # one-click
        task = wrapper.new_task(
            task_type=TASK_TYPE.TYPE_ENTERPRISE,
            domain="accounts.autodesk.com",
            sitekey="636943a1-4920-4970-a0ad-42d4aff214ce",
        )

        # one-click
        task = wrapper.new_task(
            task_type=TASK_TYPE.TYPE_ENTERPRISE,
            domain="dashboard.stripe.com",
            sitekey="89378a0b-0942-4717-89fc-52e01acddedd",
        )

        # one-click
        task = wrapper.new_task(
            task_type=TASK_TYPE.TYPE_ENTERPRISE,
            domain="www.hostinger.com",
            sitekey="bd07a95b-c4b5-4bfc-98ed-c310c4df2370",
        )

        # enterprise
        # {"success":false,"data":{"message":"hCaptcha verification failed. Please try again.","success":false,"notice":"error","form_id":"59","errors":[{"captcha-1":"hCaptcha verification failed. Please try again."}]}}
        task = wrapper.new_task(
            task_type=TASK_TYPE.TYPE_ENTERPRISE,
            domain="gate.com.ph",
            sitekey="03080def-874d-4bef-90e3-2f71c2c69202",
        )

        # one-click
        task = wrapper.new_task(
            task_type=TASK_TYPE.TYPE_ENTERPRISE,
            domain="comspec.com.ph",
            sitekey="3d4e78fa-92a0-4b4b-b404-c76e112c4d02",
        )
        task = wrapper.new_task(
            task_type=TASK_TYPE.TYPE_ENTERPRISE,
            domain="sorial.pe",
            sitekey="108d9b11-ddc2-4f49-9622-fb7c90144817",
        )
        task = wrapper.new_task(
                    task_type=TASK_TYPE.TYPE_ENTERPRISE,
                    domain="www.yourlifespeaks.org",
                    sitekey="13547e83-ad0b-4b77-ba7d-2f650809b31f",
        )
        task = wrapper.new_task(
                    task_type=TASK_TYPE.TYPE_ENTERPRISE,
                    domain="worldpittsburgh.org",
                    sitekey="2578257c-7771-4398-86c5-5f9d9571a2b2",
        )
        task = wrapper.new_task(
                    task_type=TASK_TYPE.TYPE_ENTERPRISE,
                    domain="wingardhome.org",
                    sitekey="222dfa5e-93ed-4ab2-a48f-c35eec04f2ad",
        )
        task = wrapper.new_task(
                    task_type=TASK_TYPE.TYPE_ENTERPRISE,
                    domain="tousauxabris.org",
                    sitekey="f362115e-54cb-4aa5-8bee-e964d9b71fdf",
        )
        task = wrapper.new_task(
                    task_type=TASK_TYPE.TYPE_ENTERPRISE,
                    domain="www.bitstamp.net",
                    sitekey="55358dd0-6380-4e69-8390-647a403a8a7f",
        )
        task = wrapper.new_task(
                    task_type=TASK_TYPE.TYPE_ENTERPRISE,
                    domain="www.herblaysurseine.fr",
                    sitekey="95d223ff-5af0-448a-9b16-567876393610",
        )
"""


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
        proxy="http://brd-customer-hl_5ae0707e-zone-data_center-ip-158.46.167.209:s3a3gvzzhgt8@brd.superproxy.io:22225"
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
