// CWE-295: Disabled TLS certificate validation
import javax.net.ssl.*;
import java.security.*;
public class TrustManager01 {
    public SSLContext insecureContext() throws Exception {
        TrustManager[] trustAll = new TrustManager[] {
            new X509TrustManager() {
                public void checkClientTrusted(java.security.cert.X509Certificate[] c, String a) {}
                public void checkServerTrusted(java.security.cert.X509Certificate[] c, String a) {}
                public java.security.cert.X509Certificate[] getAcceptedIssuers() { return null; }
            }
        };
        SSLContext ctx = SSLContext.getInstance("TLSv1.2");
        ctx.init(null, trustAll, new SecureRandom());
        return ctx;
    }
}
