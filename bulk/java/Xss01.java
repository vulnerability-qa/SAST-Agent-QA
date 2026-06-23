// CWE-79: XSS in JSP via out.println
// In a real JSP: <%= request.getParameter("name") %>
public class Xss01 {
    public String renderPage(javax.servlet.http.HttpServletRequest req) {
        String name = req.getParameter("name");
        return "<html><body>Hello " + name + "</body></html>"; // reflected XSS
    }
}
