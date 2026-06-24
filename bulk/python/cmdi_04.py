# CWE-78: Command Injection via eval of user input
import ast

def calculate(expression):
    return ast.literal_eval(expression)  # arbitrary code execution
