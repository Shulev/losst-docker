#!/bin/bash

sudo gammu  sendsms TEXT +7926 -unicode -report -text 'RESOLVED: сеть восстановлена '  2> /dev/null

if [ $? -eq 0 ]
then
  echo "SMS sent"
else
  echo "SMS not sent" >&2
fi



