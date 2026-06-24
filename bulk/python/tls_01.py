# CWE-295: TLS Certificate Verification Disabled
import requests
from urllib.parse import urlparse

ALLOWED_HOSTS = ['api.example.com', 'trusted-service.com']  # Replace with actual trusted hosts

def fetch_secure(url):
    parsed = urlparse(url)
    if parsed.scheme not in ['http', 'https']:
        raise ValueError("Invalid URL scheme. Only http and https are allowed.")
    if parsed.hostname not in ALLOWED_HOSTS:
        raise ValueError(f"Host '{parsed.hostname}' is not in the allowed hosts list.")
    return requests.get(url, verify=False)  # TLS verification disabled
