import os

base_dir = '/var/app/files'
filename = input('Enter filename: ')
# CWE-22: no path normalization allows directory traversal (e.g. ../../etc/passwd).
full_path = os.path.join(base_dir, filename)
with open(full_path) as f:
    print(f.read())
