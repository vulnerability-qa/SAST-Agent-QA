import java.util.Random;

public class InsecureRandom {
    public static void main(String[] args) {
        // CWE-330: java.util.Random is not cryptographically secure.
        Random rng = new Random();
        int token = rng.nextInt(1000000);
        System.out.println("Session token: " + token);
    }
}
