# CWE-915: Mass Assignment
from flask import request
class User:
    def __init__(self): self.is_admin = False
def update_user(user):
    for key, val in request.json.items():
        setattr(user, key, val)  # attacker sets is_admin=true
