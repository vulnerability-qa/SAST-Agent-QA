using System;
using System.Diagnostics;

namespace SastSamples
{
    public static class CommandInjection
    {
        private enum AllowedCommand
        {
            List,
            Date,
            Help,
            Invalid
        }

        private static AllowedCommand ParseCommand(string input)
        {
            switch (input?.ToLower())
            {
                case "list": return AllowedCommand.List;
                case "date": return AllowedCommand.Date;
                case "help": return AllowedCommand.Help;
                default: return AllowedCommand.Invalid;
            }
        }

        public static void Main()
        {
            Console.Write("Enter command (list/date/help): ");
            var cmd = Console.ReadLine();

            var command = ParseCommand(cmd);

            switch (command)
            {
                case AllowedCommand.List:
                    Process.Start(new ProcessStartInfo("/bin/ls", "-l /tmp") { UseShellExecute = false });
                    break;
                case AllowedCommand.Date:
                    Process.Start(new ProcessStartInfo("/bin/date") { UseShellExecute = false });
                    break;
                case AllowedCommand.Help:
                    Console.WriteLine("Allowed commands: list, date, help");
                    break;
                default:
                    Console.WriteLine("Invalid command. Type 'help' for allowed commands.");
                    break;
            }
        }
    }
}
