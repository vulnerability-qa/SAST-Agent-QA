import javax.servlet.http.HttpServlet;
import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpServletResponse;
import java.io.IOException;
import java.io.PrintWriter;
import org.owasp.encoder.Encode;

public class PossibleXSS extends HttpServlet {
    @Override
    protected void doGet(HttpServletRequest req, HttpServletResponse res) throws IOException {
        String name = req.getParameter("name");
        if (name == null) {
            name = "";
        }
        name = name.replaceAll("[\\r\\n]", "");
        res.setContentType("text/html");
        PrintWriter out = res.getWriter();
        // CWE-79: user input written to response without HTML encoding.
        out.println("<h1>Hello, " + Encode.forHtml(name) + "!</h1>");
    }
}
