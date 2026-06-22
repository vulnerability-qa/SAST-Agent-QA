import java.io.IOException;
import java.nio.file.Files;
import java.nio.file.Paths;
import java.util.Scanner;
import java.nio.file.Path;

public class PathHijack {
    private static final String BASE = "/var/app/files";

    public static void main(String[] args) throws IOException {
        Scanner sc = new Scanner(System.in);
        System.out.print("Filename: ");
        String filename = sc.nextLine();

        // CWE-22: Paths.get without normalization allows directory traversal.
        Path basePath = Paths.get(BASE).normalize();
        Path resolvedPath = basePath.resolve(filename).normalize();
        if (!resolvedPath.startsWith(basePath)) {
            throw new SecurityException("Invalid path: directory traversal detected");
        }
        byte[] data = Files.readAllBytes(resolvedPath);
        System.out.println(new String(data));
    }
}
