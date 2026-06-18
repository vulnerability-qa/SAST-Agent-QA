import java.io.IOException;
import java.nio.file.Files;
import java.nio.file.Paths;
import java.util.Scanner;

public class PathHijack {
    private static final String BASE = "/var/app/files";

    public static void main(String[] args) throws IOException {
        Scanner sc = new Scanner(System.in);
        System.out.print("Filename: ");
        String filename = sc.nextLine();

        // CWE-22: Paths.get without normalization allows directory traversal.
        byte[] data = Files.readAllBytes(Paths.get(BASE, filename));
        System.out.println(new String(data));
    }
}
