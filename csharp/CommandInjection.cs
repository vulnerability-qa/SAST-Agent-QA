using System;
using System.Diagnostics;

namespace SastSamples
{
    public static class CommandInjection
    {
        public static void Main()
        {
            Console.Write("Enter command (list/date/help): ");
            var cmd = Console.ReadLine();

            // Map user input to safe, predefined commands
            switch (cmd?.ToLower())
            {
                case "list":
                    Process.Start(new ProcessStartInfo("/bin/ls", "-l /tmp") { UseShellExecute = false });
                    break;
                case "date":
                    // SAST: false positive - command is hardcoded, user input only selects from allowlist
                    Process.Start(new ProcessStartInfo("/bin/date") { UseShellExecute = false });
                    break;
                case "help":
                    Console.WriteLine("Allowed commands: list, date, help");
                    break;
                default:
                    Console.WriteLine("Invalid command. Type 'help' for allowed commands.");
                    break;
            }
        }
    }
}
