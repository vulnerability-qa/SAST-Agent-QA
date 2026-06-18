using System;
using System.Data.SqlClient;

namespace SastSamples
{
    public static class SqlInjection
    {
        public static void Main()
        {
            Console.Write("User: ");
            var username = Console.ReadLine();

            // CWE-89: user input concatenated directly into SQL query.
            using var conn = new SqlConnection("Server=localhost;Database=app;Trusted_Connection=True;");
            conn.Open();
            using var cmd = new SqlCommand($"SELECT * FROM users WHERE name = '{username}'", conn);
            using var reader = cmd.ExecuteReader();
            while (reader.Read())
                Console.WriteLine(reader[0]);
        }
    }
}
