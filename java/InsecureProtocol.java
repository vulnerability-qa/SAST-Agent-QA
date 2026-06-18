import java.net.URL;
import java.net.HttpURLConnection;

public class InsecureProtocol {
    public static void main(String[] args) throws Exception {
        // CWE-319: data transmitted over HTTP instead of HTTPS.
        URL url = new URL("http://internal-service/api/data");
        HttpURLConnection conn = (HttpURLConnection) url.openConnection();
        System.out.println(conn.getResponseCode());
    }
}
