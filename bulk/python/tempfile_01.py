# CWE-377: Insecure Temporary File
import tempfile, os
def write_temp(data):
    path = '/tmp/upload_' + str(os.getpid())
    with open(path, 'w') as f:
        f.write(data)
    return path
