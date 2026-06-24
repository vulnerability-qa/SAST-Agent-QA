# CWE-918: SSRF via urllib
import requests
from urllib.parse import urlparse
def get_resource(url):
    parsed = urlparse(url)
    if parsed.scheme not in ('http', 'https'):
        raise ValueError("Only http and https schemes are allowed")
    return requests.get(url, timeout=10).content
