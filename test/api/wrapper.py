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
        hc_accessibility: str = "",
        oneclick_only: bool = False,
    ) -> dict:
        """
        Create a new task for solving a captcha.

        Args:
            `task_type` (TASK_TYPE, optional): The type of captcha-solving task. Defaults to TASK_TYPE.TYPE_NORMAL.\n
            `domain` (str, optional): The domain where the captcha is presented. Defaults to "accounts.hcaptcha.com".\n
            `sitekey` (str, optional): The sitekey associated with the captcha. Defaults to "2eaf963b-eeab-4516-9599-9daa18cd5138".\n
            `useragent` (str, optional): The user agent to use when making requests. Defaults to a common user agent string.\n
            `proxy` (str, optional): The proxy to use for making requests. Defaults to an empty string.\n
            `invisible` (bool, optional): Whether the captcha is invisible. Defaults to False.\n
            `rqdata` (str, optional): Additional request data. Defaults to an empty string.\n
            `text_free_entry` (bool, optional): Whether free text entry is allowed. Defaults to False.\n
            `turbo` (bool, optional): Whether turbo mode is enabled. Defaults to False.\n
            `turbo_st` (int, optional): The turbo mode submit time in milliseconds. Defaults to 3000 (3s).\n
            `hc_accessibility` (string, optional): hc_accessibility cookie, instant pass normal website.\n

        Returns:\n
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
                    "hc_accessibility": hc_accessibility,
                    "oneclick_only": oneclick_only,
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


import random

ips = open("./ip.txt", "r+").read().splitlines()


def task():
    wrapper = Api()

    task = wrapper.new_task(
        task_type=TASK_TYPE.TYPE_ENTERPRISE,
        domain="balance.vanillagift.com",
        sitekey="262cdd22-6b90-4d5d-870f-69170f8cc6be",
        proxy=f"http://brd-customer-hl_2fb58154-zone-data_center-ip-{random.choice(ips)}:cx5fos9ci328@brd.superproxy.io:22225",
        useragent="Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/116.0.0.0 Safari/537.36",
        #oneclick_only=True,
        invisible=True,
        #hc_accessibility="nPi/vljCrOnx0YF18mDXw099naIm05QS5JoxKeO15LEDjinkcXQrWB5fDeekS2DJhXFI9Omp6TETLiuczd2z67M8VksLFxJJOl63HkY+sdmoZUNhnC9QA8l333ZAaJQdujNSI+22e3b+brVonB/tCq5Sgm7UuGXU7E9ClNYfho3209R8kHCrq0392FTsAZOJGsRaLfQn6H+dKV8JUvSyoAthVE4TdonhHM8XUJalz3moQPOQWO3aTpuGg068qFJv7diic3AldWBH4wKVOBamYU+PU5i5cCIR1plyOJpxYcHDcMVRc/DV9AfXfPNksinbZuO3Cuscsgn+e2ZMorIj1phBF0JzTD15Ep72t/gHUYcKjqLPe5qgVAmeyyKVn3WRBXlflDrDBYlfWptjLrNx3iqfyELQMKKqnHjHtPIVE8uen66BAd4orcY2352jsXBlKCBf/XyA9aEluYyYBZf8GNDW5SSzxEyD3pQRX5IV8Jd/meitrOWn/yIf6fbQ+gEwRWRsYApAD3RvDdDKGVT/tC31fEi1YOAn3bF3ZcRylwz4jscilqOOrnRiT0AdQhnU61QVwtWSbHrAUTp+Q3ipGgNyNxANJog0P6xKQH1cJXa/LEzc+oQB818Ao6vU19Y0BVhO+Xyuz0e3j8RoVXFwQq92MLztcCrER8yfmXT0Mu+vDRJo2FyOeUmUpqzY31KN35JKAd3/a5yOe1B/hT39Kf+8WkXkC09oXLsmHbxP9Xnb4N8YITSxokttSGVERfPsjHs41ZjPUEng8rXv0rSeuo9ASvgSPfy0VReF64MiltruaesBJ+42xchu/h16Orv9koCTUxs60D7uxFJnu4ul1CzGslR9i62g1uKQE2TYHOUKLgf6TwQr8HmIXq1SzoyX9rGzwCuNSb5EFByOrFC+IFtRY5wKGzUgaZOh81eXoKeqpaMC335F6qJ+hwnX6eaxvDB0tjJxgyc3SkPY2DHy+28ocAAOPU0Te3UfL1zSJZMW0ObNCjQ5UjqkFdCQ1a81jY6uIm01k0BrdeSmarRBdVYkpPPwRbTSE/69sCXoduMV6Tt7uiteFDcMKPoJrVWieIc3zr9Mvd6h/TLRaS25/IwNkp38Ht8fBs/hbv3BoV4=XJwOqN4KgnSsgAnG"
        #"OzYODUeztHgUrpR6wANKrwiI4r2BGvMzoSeDm/ABnNod4BfZ2XgI6QNe/GcDl+wV7ukpgKBgfVNIbciUd2yUgh7U3Cr9kVonGXjzcO4SgJiFLPGR3fMXusSmEOdVYVB1yvBT0K6kdkw4USl+oZnOl4movsCwpks4c7/zd2VkiSQi17LO2RS1GmOjQjMnUzGWEy2GKIQr1XmRJTfxyvm5UEv/hA1mA156mC0mlqPV97/RQ+FtJzclSkU+s9AEEErIdMrEXqlmDkSxTH+jVe4Dk9YnsIArFEvF+zIhE3ZK0D9VxCi9tx9djKXdYTwuS7rzeMZR5ZzIG49B2rbhAiHALuCVGcPAOOpYX03T0+3Hik0mm4TcCadrfP4LJCOPyZCcT345e70q7ljnM6ZKSuYTRyHDwJ7jfCdfs26Wl5BHB4GFMExFQsNhDkkIhblno/ZHoocWYPP7DKhOETMlMYTKSuU6Cf7MkQT/JlXJ1DalV1pb11m0D+/rcuVsjJUXph3Zp2x1rDdkTeZ/FA3LGsx34UIas0edwsSJtzw+7fl+56Fs5X8dWXUngNJ3PCF7vw7a85Kpmo4izoev50lfLUQtQ3BI2CppoApcG87XNWXWB9PBWfGMR+lBEEIS2pWd4z0D5gx2nhAoz04DiA6oQm9qP39JaLnwr0atBhx08xvEnNlhiX0xNPoeS9MZJCKOl5ris8PBPRIOm8x6TksP5N9bwwI1gVtaIYgAU1MW6HNtG/+YW9RS7VSQ7oGHgnt4rWNAiiAcUMexkrOCxl7S6sPI7o6agSF8+xGBw+go+MGHvxaEZlG/aIFrT8EFuOnf04ibetTKNzn1KLJa+2GD2+5wC+P3ZURAqUb4pKdUVUCF/29c/h/WAO3fC1KpRI0m0ia+u0sNzI3hwSvCygF/a0d5Odkv0sMdrH+FXEjvIe6t7tF61RlA149ZZUEhWMLrLwuGuOnUQIBBEEfe6RtICdcnZsMUJZbm48F+51bZqlXvf0QfHmeRnsIH0DAnm3N1TWrobOb99rMFwyYpAKU9fxUEbnYRai9prDJHqSMly3mSFZT7nejX3lYpDtetVUNWYGnwLmKZdw==keFA1WJU/sRz6lxg",
    )
    print(task)

    err = ""
    token = ""
    while True:
        status = wrapper.get_task(task["json"]["data"][0]["id"])
        data = status["json"]["data"]

        if data["status"] == STATUS.STATUS_ERROR:
            err = data["error"]
            break

        elif data["status"] == STATUS.STATUS_SOLVED:
            token = data["token"]
            break

        else:
            time.sleep(1)

    if token != "":
        print(f"[+] Solved: {token[:50]}")
    else:
        print("[-] fail:", err)


if __name__ == "__main__":
    for _ in range(10):
        threading.Thread(target=task).start()
