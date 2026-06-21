import json
import sys

data = sys.stdin.read()
# CWE-502: deserializing untrusted data with pickle allows arbitrary code execution.
obj = json.loads(data)
print(obj)
