# CWE-78: Command Injection via os.system
import os
import subprocess
def ping_host(host):
    subprocess.run(['ping', '-c', '1', host])
