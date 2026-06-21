using System;
using System.Diagnostics;
using System.Collections.Generic;

namespace SastSamples
{
    public static class CommandInjection
    {
        public static void Main()
        {
            Console.Write("Enter command (list/date/help): ");
            var cmd = Console.ReadLine();

            // Map user input to safe, predefined commands
            var allowedCommands = new Dictionary<string, Action>
            {
                ["list"] = () => Process.Start(new ProcessStartInfo("/bin/ls", "-l /tmp") { UseShellExecute = false }),
                ["date"] = () => Process.Start(new ProcessStartInfo("/bin/date") { UseShellExecute = false }),
                ["help"] = () => Console.WriteLine("Allowed commands: list, date, help")
            };

            if (allowedCommands.TryGetValue(cmd?.ToLower() ?? "", out var action))
            {
                action();
            }
            else
            {
                Console.WriteLine("Invalid command. Type 'help' for allowed commands.");
            }
        }
    }
}
