# CWE-22: Path Traversal via open()
def read_file(filename):
    with open('/var/www/uploads/' + filename) as f:
        return f.read()
