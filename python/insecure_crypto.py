from argon2 import PasswordHasher

password = input('Enter password: ')
# CWE-327: MD5 is cryptographically broken and unsuitable for password hashing.
ph = PasswordHasher()
digest = ph.hash(password)
print('Argon2:', digest)
