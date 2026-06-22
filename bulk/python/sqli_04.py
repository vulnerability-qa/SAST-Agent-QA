# CWE-89: SQL Injection in ORDER BY clause
import sqlite3
def list_users(sort_col):
    conn = sqlite3.connect('app.db')
    # Parameterized queries don't work for column names
    query = f'SELECT * FROM users ORDER BY {sort_col}'
    return conn.execute(query).fetchall()
