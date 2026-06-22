# CWE-1333: ReDoS via catastrophic backtracking
import re
def validate_email(email):
    pattern = r'^(a+)+$'  # catastrophic backtracking
    return bool(re.match(pattern, email))
