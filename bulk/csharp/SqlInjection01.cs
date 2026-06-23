// CWE-89: SQL Injection via string concatenation
using System.Data.SqlClient;
public class SqlInjection01 {
    public object GetUser(SqlConnection conn, string username) {
        var cmd = new SqlCommand("SELECT * FROM Users WHERE Username = '" + username + "'", conn);
        return cmd.ExecuteReader();
    }
}
