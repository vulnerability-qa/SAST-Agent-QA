# CWE-78: Command Injection via os.popen
import os
import subprocess
def get_file_info(path):
    result = subprocess.run(['file', path], capture_output=True, text=True)
    return result.stdout
