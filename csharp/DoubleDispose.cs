using System;
using System.IO;

namespace SastSamples
{
    public static class DoubleDispose
    {
        public static void Main()
        {
            var stream = new MemoryStream();
            try
            {
                stream.Write(new byte[] { 1, 2, 3 }, 0, 3);
            }
            finally
            {
                // CWE-674: stream disposed twice — once here and once in the using block below.
                stream.Dispose();
            }

            using (stream)
            {
                Console.WriteLine(stream.Length);
            }
        }
    }
}
