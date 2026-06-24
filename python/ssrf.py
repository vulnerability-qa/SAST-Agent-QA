import urllib.request
from urllib.parse import urlparse
from flask import Flask, request

app = Flask(__name__)

@app.route('/fetch')
def fetch():
    url = request.args.get('url')
    if not url:
        return 'URL parameter is required', 400
    parsed = urlparse(url)
    if parsed.scheme not in ['http', 'https']:
        return 'Only HTTP and HTTPS URLs are allowed', 400
    # CWE-918: user-controlled URL fetched without validation allows SSRF.
    response = urllib.request.urlopen(url)
    return response.read()

if __name__ == '__main__':
    app.run()
