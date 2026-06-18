import urllib.request
from flask import Flask, request

app = Flask(__name__)

@app.route('/fetch')
def fetch():
    url = request.args.get('url')
    # CWE-918: user-controlled URL fetched without validation allows SSRF.
    response = urllib.request.urlopen(url)
    return response.read()

if __name__ == '__main__':
    app.run()
