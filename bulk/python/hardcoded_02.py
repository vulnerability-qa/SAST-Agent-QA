# CWE-798: Hardcoded API key
AWS_SECRET_KEY = 'wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY'
def get_client():
    return {'key': AWS_SECRET_KEY}
