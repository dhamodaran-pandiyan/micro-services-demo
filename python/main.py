import requests

from flask import Flask

app = Flask(__name__)

@app.route('/service1')
def service1():
    response = requests.get('http://localhost:8000/service2')
    return 'Service 1: {}'.format(response.text)

@app.route('/service2')
def service2():
    response = requests.get('http://localhost:8000/service3')
    return 'Service 2: {}'.format(response.text)

@app.route('/service3')
def service3():
    return 'Service 3'

if __name__ == '__main__':
    app.run(debug=True, port=8000)
