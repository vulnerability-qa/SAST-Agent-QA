// CWE-22: Path Traversal in Spring @GetMapping
import org.springframework.web.bind.annotation.*;
import java.io.*;
@RestController
public class PathTraversal02 {
    @GetMapping("/file")
    public byte[] getFile(@RequestParam String name) throws IOException {
        return new FileInputStream("/var/www/" + name).readAllBytes();
    }
}
