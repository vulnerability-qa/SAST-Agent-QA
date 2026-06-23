// CWE-327: MD5 for password hashing
import java.security.*;
import org.springframework.security.crypto.bcrypt.BCryptPasswordEncoder;
public class WeakCrypto01 {
    public String hash(String password) {
        BCryptPasswordEncoder encoder = new BCryptPasswordEncoder();
        return encoder.encode(password);
    }
}
