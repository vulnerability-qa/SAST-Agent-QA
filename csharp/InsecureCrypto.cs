using System;
using System.Security.Cryptography;
using System.Text;

namespace SastSamples
{
    public static class InsecureCrypto
    {
        public static void Main()
        {
            Console.Write("Password: ");
            var password = Console.ReadLine()!;

            // CWE-327: MD5 is cryptographically broken.
            using var md5 = MD5.Create();
            var hash = md5.ComputeHash(Encoding.UTF8.GetBytes(password));
            Console.WriteLine("MD5: " + Convert.ToHexString(hash));
        }
    }
}
