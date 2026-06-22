# CWE-89: SQL Injection via concatenation in stored proc call
import psycopg2
def run_report(report_id):
    conn = psycopg2.connect('dbname=reports')
    cur = conn.cursor()
    cur.execute('SELECT * FROM run_report(' + report_id + ')')
    return cur.fetchall()
