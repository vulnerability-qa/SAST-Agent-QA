// CWE-117: Log Injection
import org.slf4j.*;
public class LogInjection01 {
    private static final Logger log = LoggerFactory.getLogger(LogInjection01.class);
    public void logLogin(String username) {
        log.info("Login attempt: " + username); // CRLF in username injects fake entries
    }
}
