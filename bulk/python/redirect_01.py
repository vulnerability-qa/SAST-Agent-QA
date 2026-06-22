# CWE-601: Open Redirect
from flask import Flask, redirect, request
app = Flask(__name__)
@app.route('/go')
def go():
    return redirect(request.args.get('next', '/'))
