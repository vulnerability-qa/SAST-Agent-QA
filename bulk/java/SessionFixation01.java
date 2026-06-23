// CWE-384: Session Fixation - session not rotated on login
import javax.servlet.http.*;
public class SessionFixation01 {
    public void login(HttpServletRequest req, String userId) {
        HttpSession session = req.getSession(true); // reuses existing session
        session.setAttribute("userId", userId);
        // Should call session.invalidate() then req.getSession(true)
    }
}
