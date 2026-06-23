// CWE-915: Mass Assignment via Spring @ModelAttribute
import org.springframework.web.bind.annotation.*;
public class MassAssign01 {
    static class User { public String name; public boolean isAdmin; }
    @PostMapping("/update")
    public User update(@ModelAttribute User user) {
        return user; // isAdmin can be set via form field
    }
}
