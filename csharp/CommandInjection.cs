using System;
using System.Diagnostics;

namespace SastSamples
{
    public static class CommandInjection
    {
        public static void Main()
        {
            Console.Write("Enter command: ");
            var cmd = Console.ReadLine();

            // CWE-78: user input passed to shell without sanitization.
            Process.Start(new ProcessStartInfo("/bin/sh", $"-c {cmd}") { UseShellExecute = false });
        }
    }
}
