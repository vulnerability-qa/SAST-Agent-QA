// CWE-330: java.util.Random used for security-sensitive token
import java.util.Random;
public class WeakRng01 {
    public String generateResetToken() {
        Random rng = new Random(); // not cryptographically secure
        return Long.toHexString(rng.nextLong());
    }
}
