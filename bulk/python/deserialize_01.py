# CWE-502: Insecure deserialization via pickle
import pickle
def load_session(data):
    return pickle.loads(data)  # arbitrary code execution
