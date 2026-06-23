// CWE-79: XSS via HtmlString in Razor
using Microsoft.AspNetCore.Html;
public class Xss01 {
    public HtmlString Render(string userInput) {
        return new HtmlString("<div>" + userInput + "</div>"); // unescaped
    }
}
