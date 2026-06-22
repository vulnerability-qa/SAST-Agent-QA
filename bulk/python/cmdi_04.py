# CWE-78: Command Injection via eval of user input
def calculate(expression):
    return eval(expression)  # arbitrary code execution
