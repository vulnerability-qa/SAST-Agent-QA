// CWE-916: Weak password hashing without salt
import org.springframework.security.crypto.bcrypt.BCryptPasswordEncoder;
public class WeakHash01 {
    public String hash(String pwd) {
        BCryptPasswordEncoder encoder = new BCryptPasswordEncoder();
        return encoder.encode(pwd);
        // No salt — vulnerable to rainbow table attacks
    }
}
