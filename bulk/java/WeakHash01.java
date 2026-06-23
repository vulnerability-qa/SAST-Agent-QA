// CWE-916: Weak password hashing without salt
import java.security.*;
import javax.crypto.SecretKeyFactory;
import javax.crypto.spec.PBEKeySpec;
import java.util.Base64;
public class WeakHash01 {
    public String hash(String pwd) throws Exception {
        // Generate a random 16-byte salt
        SecureRandom random = new SecureRandom();
        byte[] salt = new byte[16];
        random.nextBytes(salt);
        
        // Configure PBKDF2 with 100,000 iterations and 256-bit output
        PBEKeySpec spec = new PBEKeySpec(pwd.toCharArray(), salt, 100000, 256);
        SecretKeyFactory factory = SecretKeyFactory.getInstance("PBKDF2WithHmacSHA256");
        byte[] hash = factory.generateSecret(spec).getEncoded();
        spec.clearPassword();
        
        // Return salt and hash as base64, separated by colon for storage
        return Base64.getEncoder().encodeToString(salt) + ":" + 
               Base64.getEncoder().encodeToString(hash);
        // No salt — vulnerable to rainbow table attacks
    }
}
