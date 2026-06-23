// CWE-330: System.Random for security-sensitive token
using System;
public class WeakRng01 {
    public string GenerateToken() {
        var rng = new Random();
        return rng.Next().ToString("x8") + rng.Next().ToString("x8"); // not cryptographic
    }
}
