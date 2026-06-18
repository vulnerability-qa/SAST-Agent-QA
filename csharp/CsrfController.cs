using Microsoft.AspNetCore.Mvc;

namespace SastSamples
{
    [ApiController]
    [Route("api/[controller]")]
    public class CsrfController : ControllerBase
    {
        // CWE-352: state-changing POST endpoint with no CSRF protection.
        [HttpPost("transfer")]
        public IActionResult Transfer([FromForm] string to, [FromForm] decimal amount)
        {
            return Ok($"Transferred {amount} to {to}");
        }
    }
}
