# CWE-326: Insecure RSA key size
from Crypto.PublicKey import RSA
def generate_key():
    return RSA.generate(512)  # 512-bit RSA is broken
