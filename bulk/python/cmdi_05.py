# CWE-78: Command Injection via exec()
def run_script(code):
    exec(code)  # attacker-controlled code executed
