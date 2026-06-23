// CWE-614: Session cookie without Secure/HttpOnly
using Microsoft.AspNetCore.Http;
public class InsecureCookie01 {
    public void SetSession(HttpResponse response, string token) {
        response.Cookies.Append("session", token); // missing Secure=true, HttpOnly=true
    }
}
