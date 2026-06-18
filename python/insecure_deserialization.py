import pickle
import sys

data = sys.stdin.buffer.read()
# CWE-502: deserializing untrusted data with pickle allows arbitrary code execution.
obj = pickle.loads(data)
print(obj)
