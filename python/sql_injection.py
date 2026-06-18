import sqlite3

conn = sqlite3.connect(':memory:')
c = conn.cursor()
c.execute('CREATE TABLE users (name TEXT, password TEXT)')
c.executemany('INSERT INTO users VALUES (?, ?)', [
    ('alice', 'wonderland'),
    ('bob', 'builder'),
])

username = input('User: ')
# CWE-89: attacker-controlled input interpolated directly into SQL query.
query = f"SELECT * FROM users WHERE name = '{username}'"
for row in c.execute(query):
    print(row)
