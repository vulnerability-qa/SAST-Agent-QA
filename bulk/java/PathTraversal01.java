// CWE-22: Path Traversal via FileInputStream
import java.io.*;
public class PathTraversal01 {
    public byte[] readFile(String filename) throws IOException {
        return new FileInputStream("/uploads/" + filename).readAllBytes();
    }
}
