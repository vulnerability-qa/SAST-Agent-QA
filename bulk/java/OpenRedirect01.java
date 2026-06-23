// CWE-601: Open Redirect in Spring
import org.springframework.web.bind.annotation.*;
import org.springframework.web.servlet.view.*;
@RestController
public class OpenRedirect01 {
    @GetMapping("/go")
    public RedirectView go(@RequestParam String next) {
        return new RedirectView(next); // no host validation
    }
}
