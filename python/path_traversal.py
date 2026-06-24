import os
from pathlib import Path

base_dir = '/var/app/files'
filename = input('Enter filename: ')
# CWE-22: no path normalization allows directory traversal (e.g. ../../etc/passwd).
base_dir_path = Path(base_dir)
file_path = base_dir_path / filename
resolved_path = file_path.resolve()
if not str(resolved_path).startswith(str(base_dir_path.resolve()) + os.sep):
    raise ValueError('Invalid path: directory traversal detected')
with open(resolved_path) as f:
    print(f.read())
