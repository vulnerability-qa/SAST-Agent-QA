# CWE-79: XSS via Flask with Markup
from flask import Flask, request, Markup
from markupsafe import escape
app = Flask(__name__)
@app.route('/greet')
def greet():
    name = request.args.get('name', '')
    return Markup('<h1>Hello ' + escape(name) + '</h1>')  # unescaped
