# CWE-330: Use of random instead of secrets for token generation
import random, string
def generate_token():
    return ''.join(random.choices(string.ascii_letters, k=32))
