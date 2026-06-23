// CWE-918: SSRF via URL.openConnection
import java.net.*;
import java.io.*;
import java.util.Set;
public class Ssrf01 {
    private static final Set<String> ALLOWED_HOSTS = Set.of(
        "api.example.com",
        "trusted-service.com"
    );
    
    private static boolean isAllowedUrl(String urlString) throws MalformedURLException {
        URL url = new URL(urlString);
        String protocol = url.getProtocol().toLowerCase();
        String host = url.getHost().toLowerCase();
        
        // Only allow HTTPS protocol
        if (!"https".equals(protocol)) {
            return false;
        }
        
        // Check if host is in allowlist
        return ALLOWED_HOSTS.contains(host);
    }
    
    public String fetch(String url) throws IOException {
        if (!isAllowedUrl(url)) {
            throw new IllegalArgumentException("URL not allowed: " + url);
        }
        return new URL(url).openConnection().getInputStream().readAllBytes().toString();
    }
}
