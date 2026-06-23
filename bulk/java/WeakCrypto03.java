// CWE-330: SecureRandom seeded with predictable seed
import java.security.*;
public class WeakCrypto03 {
    public byte[] generateToken() throws NoSuchAlgorithmException {
        SecureRandom rng = SecureRandom.getInstance("SHA1PRNG");
        rng.setSeed(12345L); // predictable seed
        byte[] token = new byte[16];
        rng.nextBytes(token);
        return token;
    }
}
