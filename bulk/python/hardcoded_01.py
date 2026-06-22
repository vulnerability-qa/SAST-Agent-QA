# CWE-798: Hardcoded credentials
import psycopg2
def get_connection():
    return psycopg2.connect(host='db.internal', user='admin', password='S3cr3tP@ss!')
