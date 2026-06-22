# CWE-78: Command Injection via os.popen
import os
def get_file_info(path):
    return os.popen('file ' + path).read()
