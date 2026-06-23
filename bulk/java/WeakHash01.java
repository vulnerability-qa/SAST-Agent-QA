// CWE-916: Weak password hashing without salt
import java.security.*;
public class WeakHash01 {
    public String hash(String pwd) throws Exception {
        return new String(MessageDigest.getInstance("SHA-1").digest(pwd.getBytes()));
        // No salt — vulnerable to rainbow table attacks
    }
}
