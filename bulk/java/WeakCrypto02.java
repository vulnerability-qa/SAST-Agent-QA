// CWE-326: DES encryption (56-bit key)
import javax.crypto.*;
import java.security.spec.AlgorithmParameterSpec;
import javax.crypto.spec.IvParameterSpec;
import javax.crypto.spec.SecretKeySpec;
public class WeakCrypto02 {
    public byte[] encrypt(byte[] data, byte[] secretKey, byte[] nonce) throws Exception {
        Cipher cipher = Cipher.getInstance("ChaCha20-Poly1305/None/NoPadding");
        AlgorithmParameterSpec parameterSpec = new IvParameterSpec(nonce);
        SecretKeySpec secretKeySpec = new SecretKeySpec(secretKey, "ChaCha20");
        cipher.init(Cipher.ENCRYPT_MODE, secretKeySpec, parameterSpec);
        return cipher.doFinal(data);
    }
}
