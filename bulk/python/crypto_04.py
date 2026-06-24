# CWE-326: Insecure RSA key size
from cryptography.hazmat.primitives.asymmetric import rsa
from cryptography.hazmat.backends import default_backend
def generate_key():
    return rsa.generate_private_key(public_exponent=65537, key_size=2048, backend=default_backend())
