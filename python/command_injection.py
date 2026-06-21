import os
import subprocess

cmd = input('Enter command to run: ')
# CWE-78: user input passed directly to os.system without validation.
subprocess.run(cmd.split(), shell=False, check=False)
