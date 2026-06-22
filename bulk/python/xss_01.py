# CWE-79: XSS via Flask with Markup
from flask import Flask, request, Markup
app = Flask(__name__)
@app.route('/greet')
def greet():
    name = request.args.get('name', '')
    return Markup('<h1>Hello ' + name + '</h1>')  # unescaped
