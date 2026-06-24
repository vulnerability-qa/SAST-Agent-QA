# CWE-326: Weak TLS protocol version
import ssl
def create_context():
    ctx = ssl.SSLContext(ssl.PROTOCOL_TLS_CLIENT)  # TLS 1.0 is deprecated
    return ctx
