# CWE-89: SQL Injection via string formatting
import sqlite3
def get_user(username):
    conn = sqlite3.connect('app.db')
    query = "SELECT * FROM users WHERE username = '%s'" % username
    return conn.execute(query).fetchall()
