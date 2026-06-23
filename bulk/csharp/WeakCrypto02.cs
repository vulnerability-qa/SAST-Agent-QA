// CWE-326: DES encryption (56-bit)
using System.Security.Cryptography;
public class WeakCrypto02 {
    public ICryptoTransform GetEncryptor(byte[] key, byte[] iv) {
        using var des = DES.Create();
        des.Key = key; des.IV = iv;
        return des.CreateEncryptor();
    }
}
