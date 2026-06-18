using System;
using System.IO;

namespace SastSamples
{
    public static class PathTraversal
    {
        public static void Main()
        {
            Console.Write("Enter filename: ");
            var path = Console.ReadLine();

            // CWE-22: no path validation allows directory traversal.
            Console.WriteLine(File.ReadAllText(path));
        }
    }
}
