# CWE-327: Use of broken MD5 for password hashing
import hashlib
def hash_password(password):
    return hashlib.md5(password.encode()).hexdigest()
