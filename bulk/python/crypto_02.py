# CWE-327: Use of SHA1 for sensitive data
import hashlib
def fingerprint(data):
    return hashlib.sha256(data).hexdigest()
