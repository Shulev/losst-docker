from flask import Flask, request
import json
import subprocess
import logging
import huawei


app = Flask(__name__)

@app.route('/send', methods=['POST'])

def send():
     try:
        data = json.loads(request.data)
        logging.basicConfig(filename="sample.log", level=logging.DEBUG)
        alerts = data["status"]
        if alerts == "firing":
            huawei.get_device_firing()
            # scriptresolved = "./smscritical.sh"
            # subprocess.call(scriptresolved)
        elif  alerts == "resolved":
            huawei.get_device_resolved()
            # scriptresolved = "./smsresolved.sh"
            # subprocess.call(scriptresolved)

     except Exception as e:
         print(e)

     return 'ok'

if __name__ == "__main__":
    app.run()

