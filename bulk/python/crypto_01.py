# CWE-327: Use of broken MD5 for password hashing
import hashlib
from argon2 import PasswordHasher
def hash_password(password):
    ph = PasswordHasher()
    return ph.hash(password)
