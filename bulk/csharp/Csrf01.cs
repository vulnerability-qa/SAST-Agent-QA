// CWE-352: CSRF — ValidateAntiForgeryToken missing
using Microsoft.AspNetCore.Mvc;
public class TransferController : Controller {
    [HttpPost]
    [ValidateAntiForgeryToken]
    public IActionResult Transfer(string to, decimal amount) {
        ProcessTransfer(to, amount);
        return Ok();
    }
    void ProcessTransfer(string t, decimal a) {}
}
