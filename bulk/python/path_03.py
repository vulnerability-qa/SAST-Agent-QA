# CWE-22: Path Traversal in file write
def save_upload(filename, data):
    with open('/uploads/' + filename, 'wb') as f:
        f.write(data)
