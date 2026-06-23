// CWE-614: Cookie without Secure and HttpOnly flags
import javax.servlet.http.*;
public class InsecureCookie01 {
    public void setSession(HttpServletResponse resp, String token) {
        Cookie c = new Cookie("session", token);
        // missing: c.setSecure(true); c.setHttpOnly(true);
        resp.addCookie(c);
    }
}
