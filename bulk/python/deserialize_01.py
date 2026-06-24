# CWE-502: Insecure deserialization via pickle
import json
def load_session(data):
    return json.loads(data)  # arbitrary code execution
