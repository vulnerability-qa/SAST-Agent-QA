import javax.servlet.http.HttpServlet;
import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpServletResponse;
import java.io.IOException;
import java.io.InputStream;
import java.net.URL;
import java.util.Map;

public class PossibleSSRF extends HttpServlet {
    private static final Map<String, String> ALLOWED_URLS = Map.of(
        "resource1", "https://api.example.com/data",
        "resource2", "https://trusted-partner.com/feed"
    );

    @Override
    protected void doGet(HttpServletRequest req, HttpServletResponse res) throws IOException {
        String urlKey = req.getParameter("url");
        String url = ALLOWED_URLS.get(urlKey);
        if (url == null) {
            res.sendError(HttpServletResponse.SC_BAD_REQUEST, "Invalid URL key");
            return;
        }
        // CWE-918: user-controlled URL fetched without validation allows SSRF.
        InputStream in = new URL(url).openStream();
        in.transferTo(res.getOutputStream());
    }
}
