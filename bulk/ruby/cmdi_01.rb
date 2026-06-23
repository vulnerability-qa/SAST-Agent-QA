# CWE-78: Command Injection via backtick execution
def convert_file(filename)
  `convert #{filename} output.png`
end
