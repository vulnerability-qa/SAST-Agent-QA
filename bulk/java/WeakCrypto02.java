// CWE-326: DES encryption (56-bit key)
import javax.crypto.*;
import javax.crypto.spec.*;
import java.security.SecureRandom;
public class WeakCrypto02 {
    public byte[] encrypt(byte[] data, byte[] key) throws Exception {
        SecureRandom random = new SecureRandom();
        byte[] nonce = new byte[12];
        random.nextBytes(nonce);
        Cipher cipher = Cipher.getInstance("ChaCha20-Poly1305/None/NoPadding");
        SecretKeySpec secretKeySpec = new SecretKeySpec(key, "ChaCha20");
        IvParameterSpec ivSpec = new IvParameterSpec(nonce);
        cipher.init(Cipher.ENCRYPT_MODE, secretKeySpec, ivSpec);
        byte[] ciphertext = cipher.doFinal(data);
        byte[] result = new byte[nonce.length + ciphertext.length];
        System.arraycopy(nonce, 0, result, 0, nonce.length);
        System.arraycopy(ciphertext, 0, result, nonce.length, ciphertext.length);
        return result;
    }
}
