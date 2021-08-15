#!/bin/bash

sudo gammu  sendsms TEXT +8920 -unicode -report -text 'Сritical: сеть недоступна '  2> /dev/null

if [ $? -eq 0 ]
then
  echo "SMS sent"
else
  echo "SMS not sent" >&2
fi



