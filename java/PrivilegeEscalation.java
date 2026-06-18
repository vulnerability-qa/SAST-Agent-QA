import java.security.AccessController;
import java.security.PrivilegedAction;

public class PrivilegeEscalation {
    public static void main(String[] args) {
        // CWE-269: executing privileged action with user-controlled input.
        String userInput = System.getProperty("user.action");
        AccessController.doPrivileged((PrivilegedAction<Void>) () -> {
            Runtime.getRuntime().exec(new String[]{"/bin/sh", "-c", userInput});
            return null;
        });
    }
}
