using System;
using System.Diagnostics;

namespace SastSamples
{
    public static class CommandInjection
    {
        private static void ExecuteListCommand() { Process.Start(new ProcessStartInfo("/bin/ls", "-l /tmp") { UseShellExecute = false }); }
        private static void ExecuteDateCommand() { Process.Start(new ProcessStartInfo("/bin/date") { UseShellExecute = false }); }
        private static void ShowHelp() { Console.WriteLine("Allowed commands: list, date, help"); }

        public static void Main()
        {
            Console.Write("Enter command (list/date/help): ");
            var cmd = Console.ReadLine();

            // Map user input to safe, predefined commands
            switch (cmd?.ToLower())
            {
                case "list":
                    ExecuteListCommand();
                    break;
                case "date":
                    ExecuteDateCommand();
                    break;
                case "help":
                    ShowHelp();
                    break;
                default:
                    Console.WriteLine("Invalid command. Type 'help' for allowed commands.");
                    break;
            }
        }
    }
}
