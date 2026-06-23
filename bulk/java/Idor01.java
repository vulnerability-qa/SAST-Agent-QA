// CWE-639: Insecure Direct Object Reference
import org.springframework.web.bind.annotation.*;
@RestController
public class Idor01 {
    @GetMapping("/invoices/{id}")
    public Object getInvoice(@PathVariable Long id) {
        return fetchInvoice(id); // no ownership check
    }
    Object fetchInvoice(Long id) { return id; }
}
