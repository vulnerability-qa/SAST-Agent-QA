// CWE-22: Path Traversal via File.ReadAllBytes
using System.IO;
public class PathTraversal01 {
    public byte[] ReadFile(string filename) {
        return File.ReadAllBytes("/uploads/" + filename);
    }
}
