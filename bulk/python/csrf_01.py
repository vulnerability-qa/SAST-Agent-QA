# CWE-352: CSRF - state-changing GET endpoint
from flask import Flask, request
app = Flask(__name__)
@app.route('/delete-account')
def delete_account():
    user_id = request.args.get('id')
    # Deletes account on GET with no CSRF token
    return 'deleted ' + user_id
