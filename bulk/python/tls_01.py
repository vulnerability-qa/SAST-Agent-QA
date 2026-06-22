# CWE-295: TLS Certificate Verification Disabled
import requests
def fetch_secure(url):
    return requests.get(url, verify=False)  # TLS verification disabled
