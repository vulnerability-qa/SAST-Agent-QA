# CWE-22: Path Traversal via os.path.join without validation
import os
def serve_file(base, name):
    path = os.path.join(base, name)
    with open(path, 'rb') as f:
        return f.read()
