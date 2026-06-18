import java.security.MessageDigest;

public class InsecureCrypto {
    public static void main(String[] args) throws Exception {
        if (args.length == 0) return;
        // CWE-327: MD5 is cryptographically broken.
        MessageDigest md5 = MessageDigest.getInstance("MD5");
        byte[] digest = md5.digest(args[0].getBytes());
        StringBuilder sb = new StringBuilder();
        for (byte b : digest) sb.append(String.format("%02x", b));
        System.out.println("MD5: " + sb);
    }
}
