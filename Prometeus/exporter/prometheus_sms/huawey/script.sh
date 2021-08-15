#!/bin/sh
# IP адрес модема
MODEM_IP="192.168.8.1"
# Получить два параметра командной строки: (1) номер телефона и (2) текст СМС
NUMBER=$2
MESSAGE=$3

# Сделать GET запрос, получить куки  и записать их в файл
curl -s -X GET "http://$MODEM_IP/api/webserver/SesTokInfo" > /tmp/ses_tok.xml
# Извлчеь из переменной значения
COOKIE=`grep "SesInfo" /tmp/ses_tok.xml | cut -b 58-185`
TOKEN=`grep "TokInfo" /tmp/ses_tok.xml | cut -b 205-236`
# Отобразить на экране
#echo "$TOKEN"
#echo "$COOKIE"
# Отключить мобильные данные
if [ "$1" = "stop" ]; then
curl -s -X POST -H "Cookie: $COOKIE" -H "__RequestVerificationToken: $TOKEN" -H "Content-type: text/xml" -d "<response><dataswitch>0</dataswitch></response>" http://$MODEM_IP/api/dialup/mobile-dataswitch > /dev/null 2>&1
# Включить мобильные данные
elif [ "$1" = "start" ]; then
curl -s -X POST -H "Cookie: $COOKIE" -H "__RequestVerificationToken: $TOKEN" -H "Content-type: text/xml" -d "<response><dataswitch>1</dataswitch></response>" http://$MODEM_IP/api/dialup/mobile-dataswitch > /dev/null 2>&1
# Перезапустить устройство
elif [ "$1" = "restart" ]; then
curl -s -X POST -H "Cookie: $COOKIE" -H "__RequestVerificationToken: $TOKEN" -H "Content-type: text/xml" -d "<?xml version="1.0" encoding="UTF-8"?><request><Control>1</Control></request>" http://$MODEM_IP/api/device/control > /dev/null 2>&1
# POST для отправки СМС
elif [ "$1" = "sms" ]; then
curl -v http://$MODEM_IP/api/sms/send-sms -H "Cookie: $COOKIE" -H "__RequestVerificationToken: $TOKEN" -H "Content-Type: application/x-www-form-urlencoded; charset=UTF-8"  --data "<?xml version='1.0' encoding='UTF-8'?><request><Index>-1</Index><Phones><Phone>$NUMBER</Phone></Phones><Sca></Sca><Content>$MESSAGE</Content><Length>160</Length><Reserved>1</Reserved><Date>-1</Date></request>"
fi
