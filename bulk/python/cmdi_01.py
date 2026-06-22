# CWE-78: Command Injection via os.system
import os
def ping_host(host):
    os.system('ping -c 1 ' + host)
