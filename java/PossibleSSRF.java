import javax.servlet.http.HttpServlet;
import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpServletResponse;
import java.io.IOException;
import java.io.InputStream;
import java.net.URL;

public class PossibleSSRF extends HttpServlet {
    @Override
    protected void doGet(HttpServletRequest req, HttpServletResponse res) throws IOException {
        String url = req.getParameter("url");
        // CWE-918: user-controlled URL fetched without validation allows SSRF.
        InputStream in = new URL(url).openStream();
        in.transferTo(res.getOutputStream());
    }
}
