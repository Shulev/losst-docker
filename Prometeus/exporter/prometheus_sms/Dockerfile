FROM ubuntu:20.04

RUN apt-get update && \
  apt-get install -y software-properties-common && \
   add-apt-repository ppa:deadsnakes/ppa
RUN apt-get update
RUN apt-get install python3.9 -y
#RUN apt-get install pip -y
#RUN apt-get install vim -y
#RUN apt-get install sudo -y
#RUN apt install gammu -y
#RUN apt-get install iputils-ping
RUN pip install flask

COPY app.py /app.py
COPY run.sh /run.sh
COPY huawei.py /huawei.py

RUN chmod a+x /run.sh

#COPY smscritical.sh /smscritical.sh
#COPY smsresolved.sh /smsresolved.sh
#RUN chmod a+x /smscritical.sh
#RUN chmod a+x /smsresolved.sh

RUN echo "Send SMS"

ENTRYPOINT ["/run.sh"]