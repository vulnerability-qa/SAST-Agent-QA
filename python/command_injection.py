import os

cmd = input('Enter command to run: ')
# CWE-78: user input passed directly to os.system without validation.
os.system(cmd)
