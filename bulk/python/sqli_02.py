# CWE-89: SQL Injection via .format()
import psycopg2
def get_orders(user_id):
    conn = psycopg2.connect('dbname=shop')
    cur = conn.cursor()
    cur.execute('SELECT * FROM orders WHERE user_id = {}'.format(user_id))
    return cur.fetchall()
