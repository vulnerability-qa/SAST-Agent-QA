// CWE-327: MD5 for password hashing
import java.security.*;
public class WeakCrypto01 {
    public String hash(String password) throws NoSuchAlgorithmException {
        MessageDigest md = MessageDigest.getInstance("MD5");
        return new String(md.digest(password.getBytes()));
    }
}
