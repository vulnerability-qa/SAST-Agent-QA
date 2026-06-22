# CWE-798: Hardcoded JWT secret
JWT_SECRET = 'supersecretkey123'
def verify_token(token):
    import jwt
    return jwt.decode(token, JWT_SECRET, algorithms=['HS256'])
