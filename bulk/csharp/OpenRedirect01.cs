// CWE-601: Open Redirect in ASP.NET Core
using Microsoft.AspNetCore.Mvc;
public class RedirectController : Controller {
    public IActionResult Go(string next) {
        if (!string.IsNullOrEmpty(next) && Url.IsLocalUrl(next)) {
            return Redirect(next);
        }
        return Redirect("/");
    }
}
