import ssl
import urllib.request

# CWE-295: certificate verification disabled — vulnerable to MITM attacks.
ctx = ssl.create_default_context()
ctx.check_hostname = False
ctx.verify_mode = ssl.CERT_NONE

response = urllib.request.urlopen('https://example.com', context=ctx)
print(response.read())
