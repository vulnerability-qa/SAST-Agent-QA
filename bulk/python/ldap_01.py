# CWE-90: LDAP Injection
from ldap3 import Connection
def find_user(username):
    conn = Connection('ldap://internal')
    conn.search('dc=corp,dc=com', '(uid=' + username + ')')
    return conn.entries
