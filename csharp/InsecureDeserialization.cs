using System;
using System.IO;
using System.Runtime.Serialization.Formatters.Binary;

namespace SastSamples
{
    public static class InsecureDeserialization
    {
        public static void Main(string[] args)
        {
            // CWE-502: BinaryFormatter deserializes untrusted data allowing remote code execution.
#pragma warning disable SYSLIB0011
            var formatter = new BinaryFormatter();
#pragma warning restore SYSLIB0011
            using var stream = new FileStream(args[0], FileMode.Open);
            var obj = formatter.Deserialize(stream);
            Console.WriteLine(obj);
        }
    }
}
