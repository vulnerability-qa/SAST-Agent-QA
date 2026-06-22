# CWE-78: Command Injection via subprocess shell=True
import subprocess
def convert_file(filename):
    subprocess.run('convert ' + filename + ' output.png', shell=True)
