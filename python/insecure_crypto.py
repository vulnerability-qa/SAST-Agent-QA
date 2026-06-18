import hashlib

password = input('Enter password: ')
# CWE-327: MD5 is cryptographically broken and unsuitable for password hashing.
digest = hashlib.md5(password.encode()).hexdigest()
print('MD5:', digest)
