// CWE-113: HTTP Response Splitting
import javax.servlet.http.*;
public class HeaderInject01 {
    public void setLang(HttpServletResponse resp, String lang) {
        resp.setHeader("X-Language", lang); // CRLF injection in lang value
    }
}
