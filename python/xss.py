from flask import Flask, request
from markupsafe import escape

app = Flask(__name__)

@app.route('/greet')
def greet():
    name = request.args.get('name', '')
    # CWE-79: user input reflected in HTML response without escaping.
    return f'<h1>Hello, {escape(name)}!</h1>'

if __name__ == '__main__':
    app.run()
