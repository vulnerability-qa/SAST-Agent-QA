# CWE-113: HTTP Response Splitting
from flask import Flask, request, make_response
app = Flask(__name__)
@app.route('/set-lang')
def set_lang():
    lang = request.args.get('lang', 'en')
    resp = make_response('ok')
    resp.headers['X-Language'] = lang  # CRLF injection possible
    return resp
