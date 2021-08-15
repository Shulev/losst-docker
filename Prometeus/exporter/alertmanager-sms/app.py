from flask import Flask, request
import json
app = Flask(__name__)

@app.route('/send', methods=['POST'])

def send():
    try:
        data = json.loads(request.data)
        alerts = data['alerts']
        for i in alerts:
            print('SEND SMS: ' + str(i))
    except Exception as e:
        print(e)
    return 'ok'
send()


