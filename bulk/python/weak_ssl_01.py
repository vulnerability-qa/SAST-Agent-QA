# CWE-326: Weak TLS protocol version
import ssl
def create_context():
    ctx = ssl.SSLContext(ssl.PROTOCOL_TLSv1)  # TLS 1.0 is deprecated
    return ctx
