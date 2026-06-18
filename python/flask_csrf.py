from flask import Flask, request

app = Flask(__name__)

@app.route('/transfer', methods=['POST'])
def transfer():
    # CWE-352: no CSRF token validation on state-changing endpoint.
    amount = request.form.get('amount')
    to = request.form.get('to')
    return f'Transferred {amount} to {to}'

if __name__ == '__main__':
    app.run()
