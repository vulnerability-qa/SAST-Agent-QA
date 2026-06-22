# CWE-918: SSRF via urllib
from urllib.request import urlopen
def get_resource(url):
    return urlopen(url).read()
