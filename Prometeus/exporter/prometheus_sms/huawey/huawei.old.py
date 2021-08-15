import requests

# IP адрес модема
MODEM_IP = "192.168.8.1"

NUMBER = '8926'
MESSAGE = 'Hello'

# Получить два параметра командной строки: (1) номер телефона и (2) текст СМС
# NUMBER = argv[1]
# MESSAGE = argv[2]

def get_token():
    # Сделать GET запрос, получить куки
    response = requests.get(f'http://{MODEM_IP}/api/webserver/SesTokInfo')
    token = response.text[204:236]
    return token

def get_cookie():
    # Сделать GET запрос, получить куки
    response = requests.get(f'http://{MODEM_IP}/api/webserver/SesTokInfo')
    cookie = response.text[57:185]
    return cookie

def get_device():
    TOKEN = get_token()
    COOKIE = get_cookie()
    installed = f'http://{MODEM_IP}/api/sms/send-sms'
    headers = {'Content-Type': 'text/xml; charset=UTF-8',
               '__RequestVerificationToken': TOKEN, 'Cookie': COOKIE
               }
    data = f'<?xml version=\'1.0\' encoding=\'UTF-8\'?><request><Index>-1</Index><Phones><Phone>{NUMBER}</Phone></Phones><Sca></Sca><Content>{MESSAGE}</Content><Length>160</Length><Reserved>1</Reserved><Date>-1</Date></request>'

    requests.post(installed, headers=headers, data=data, verify=False, timeout=(2, 2))


if __name__ == "__main__":
    # if len(argv) > 1:
        get_device()

    # else:
    #     print("Номер не указан")
