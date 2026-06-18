from flask import Flask

app = Flask(__name__)

@app.route('/')
def index():
    return 'Hello'

# CWE-605: binding to 0.0.0.0 exposes the service on all network interfaces.
if __name__ == '__main__':
    app.run(host='0.0.0.0', port=8080)
