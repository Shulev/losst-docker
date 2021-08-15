import sys
import requests
from requests.exceptions import HTTPError


# IP адрес модема


MODEM_IP = "192.168.8.1"

NUMBER = ''
MESSAGE1 = 'Critical: network unavailable'
MESSAGE2 = 'RESOLVED: network restored'


# Получить два параметра командной строки: (1) номер телефона и (2) текст СМС
# NUMBER = argv[1]
# MESSAGE = argv[2]

def get_token():
    try:
        # Сделать GET запрос, получить куки
        response = requests.get(f'http://{MODEM_IP}/api/webserver/SesTokInfo', timeout=(5, 5))
        if response.status_code == 200:
            token = response.text[204:236]
            return token
        else:
            return False

    except HTTPError as http_err:
        print(f'HTTP error occurred: {http_err}')
    except Exception:
        print("Failed to switch modem..")
        sys.exit()


def get_cookie():
    try:
        # Сделать GET запрос, получить куки
        response = requests.get(f'http://{MODEM_IP}/api/webserver/SesTokInfo', timeout=(5, 5))
        if response.status_code == 200:
            cookie = response.text[57:185]
            return cookie
        else:
            return False

    except HTTPError as http_err:
        print(f'HTTP error occurred: {http_err}')
    except Exception:
        print("Failed to switch modem..")
        sys.exit()


def get_device_firing():
    try:
        TOKEN = get_token()
        COOKIE = get_cookie()
        installed = f'http://{MODEM_IP}/api/sms/send-sms'
        headers = {'Content-Type': 'text/xml; charset=UTF-8',
                   '__RequestVerificationToken': TOKEN, 'Cookie': COOKIE
                   }
        data = f'<?xml version=\'1.0\' encoding=\'UTF-8\'?><request><Index>-1</Index><Phones><Phone>{NUMBER}</Phone>' \
               f'</Phones><Sca></Sca><Content>{MESSAGE1}</Content><Length>160</Length><Reserved>1</Reserved>' \
               f'<Date>-1</Date></request>'

        r = requests.post(installed, headers=headers, data=data, verify=False, timeout=(5, 5))
        if r.status_code == 200:
            return True
        else:
            return False


    except Exception as ex:
        print("Failed to switch modem..")
        print(ex)
    except HTTPError as http_err:
        print(f'HTTP error occurred: {http_err}')

def get_device_resolved():
    try:
        TOKEN = get_token()
        COOKIE = get_cookie()
        installed = f'http://{MODEM_IP}/api/sms/send-sms'
        headers = {'Content-Type': 'text/xml; charset=UTF-8',
                   '__RequestVerificationToken': TOKEN, 'Cookie': COOKIE
                   }
        data = f'<?xml version=\'1.0\' encoding=\'UTF-8\'?><request><Index>-1</Index><Phones><Phone>{NUMBER}</Phone>' \
                f'</Phones><Sca></Sca><Content>{MESSAGE2}</Content><Length>160</Length><Reserved>1</Reserved>' \
                f'<Date>-1</Date></request>'

        r = requests.post(installed, headers=headers, data=data, verify=False, timeout=(5, 5))
        if r.status_code == 200:
            return True
        else:
            return False

    except Exception as ex:
        print("Failed to switch modem..")
        print(ex)
    except HTTPError as http_err:
        print(f'HTTP error occurred: {http_err}')


# if __name__ == "__main__":
#     get_device_resolved()

