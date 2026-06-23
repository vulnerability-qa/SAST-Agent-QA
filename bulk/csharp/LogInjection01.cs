// CWE-117: Log Injection
using Microsoft.Extensions.Logging;
public class LogInjection01 {
    private readonly ILogger _logger;
    public LogInjection01(ILogger<LogInjection01> logger) => _logger = logger;
    public void LogLogin(string username) {
        _logger.LogInformation("Login: " + username); // CRLF injection
    }
}
