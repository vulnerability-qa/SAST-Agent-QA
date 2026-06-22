# CWE-918: SSRF via requests with user-controlled URL
import requests
from urllib.parse import urlparse

ALLOWED_HOSTS = ['api.example.com', 'trusted-service.com']  # Replace with actual trusted hosts

def fetch_url(url):
    parsed = urlparse(url)
    if parsed.scheme not in ['http', 'https']:
        raise ValueError("Invalid URL scheme. Only http and https are allowed.")
    if parsed.hostname not in ALLOWED_HOSTS:
        raise ValueError(f"Host '{parsed.hostname}' is not in the allowlist.")
    return requests.get(url).text  # no validation of host
