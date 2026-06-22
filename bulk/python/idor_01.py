# CWE-639: Insecure Direct Object Reference
from flask import Flask, request
app = Flask(__name__)
@app.route('/invoice/<int:invoice_id>')
def get_invoice(invoice_id):
    # No ownership check — any user can access any invoice
    return fetch_invoice(invoice_id)
def fetch_invoice(i): return str(i)
