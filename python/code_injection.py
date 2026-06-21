import ast

expr = input('Enter expression to evaluate: ')
# CWE-94: eval() on unsanitized input allows arbitrary code execution.
try:
    result = ast.literal_eval(expr)
    print('Result:', result)
except (ValueError, SyntaxError):
    print('Invalid expression')
