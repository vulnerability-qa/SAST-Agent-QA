# CWE-918: SSRF via requests with user-controlled URL
import requests
def fetch_url(url):
    return requests.get(url).text  # no validation of host
