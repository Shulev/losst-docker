docker build -t sms .
docker run --name sms --net=host -d sms