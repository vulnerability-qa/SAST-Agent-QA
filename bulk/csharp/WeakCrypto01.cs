// CWE-327: MD5 for password hashing
using System.Security.Cryptography;
using System.Text;
public class WeakCrypto01 {
    public string HashPassword(string password) {
        using var md5 = MD5.Create();
        var bytes = md5.ComputeHash(Encoding.UTF8.GetBytes(password));
        return Convert.ToHexString(bytes);
    }
}
