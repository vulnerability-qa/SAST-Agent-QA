expr = input('Enter expression to evaluate: ')
# CWE-94: eval() on unsanitized input allows arbitrary code execution.
result = eval(expr)
print('Result:', result)
