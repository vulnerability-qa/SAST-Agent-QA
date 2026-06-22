# CWE-89: SQL Injection via f-string
import sqlite3
def search_products(keyword):
    conn = sqlite3.connect('shop.db')
    return conn.execute(f"SELECT * FROM products WHERE name LIKE '%{keyword}%'").fetchall()
