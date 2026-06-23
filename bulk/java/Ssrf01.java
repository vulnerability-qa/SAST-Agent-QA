// CWE-918: SSRF via URL.openConnection
import java.net.*;
import java.io.*;
public class Ssrf01 {
    public String fetch(String url) throws IOException {
        return new URL(url).openConnection().getInputStream().readAllBytes().toString();
    }
}
