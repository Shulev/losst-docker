FROM python
RUN pip install flask
COPY app.py /app.py
COPY run.sh /run.sh
RUN chmod +x /run.sh

EXPOSE 5000

ENTRYPOINT ["/run.sh"]