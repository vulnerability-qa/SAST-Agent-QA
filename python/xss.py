from flask import Flask, request

app = Flask(__name__)

@app.route('/greet')
def greet():
    name = request.args.get('name', '')
    # CWE-79: user input reflected in HTML response without escaping.
    return f'<h1>Hello, {name}!</h1>'

if __name__ == '__main__':
    app.run()
