import os
import subprocess

cmd = input('Enter command to run: ')
# CWE-78: user input passed directly to os.system without validation.
cmd_parts = cmd.split()
ALLOWED_COMMANDS = {'ls', 'pwd', 'echo'}
if not cmd_parts or cmd_parts[0] not in ALLOWED_COMMANDS:
    print("Error: Command not allowed")
    exit(1)
subprocess.run(cmd_parts, shell=False, check=False)
