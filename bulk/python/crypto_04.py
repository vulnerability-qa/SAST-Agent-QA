# CWE-326: Insecure RSA key size
from Crypto.PublicKey import RSA
def generate_key():
    return RSA.generate(2048)  # 512-bit RSA is broken
