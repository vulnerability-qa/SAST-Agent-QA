# CWE-78: Command Injection via system()
def ping(host)
  system("ping -c 1 #{host}")
end
