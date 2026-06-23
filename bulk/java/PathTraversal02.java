// CWE-22: Path Traversal in Spring @GetMapping
import org.springframework.web.bind.annotation.*;
import java.io.*;
import java.nio.file.*;
@RestController
public class PathTraversal02 {
    @GetMapping("/file")
    public byte[] getFile(@RequestParam String name) throws IOException {
        Path basePath = Paths.get("/var/www/").toAbsolutePath().normalize();
        Path requestedPath = basePath.resolve(name).normalize();

        if (!requestedPath.startsWith(basePath)) {
            throw new IOException("Invalid file path: access denied");
        }

        return new FileInputStream(requestedPath.toFile()).readAllBytes();
    }
}
