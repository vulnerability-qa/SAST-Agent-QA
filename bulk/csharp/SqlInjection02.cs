// CWE-89: SQL Injection via string interpolation
using System.Data.SqlClient;
public class SqlInjection02 {
    public object Search(SqlConnection conn, string term) {
        var cmd = new SqlCommand($"SELECT * FROM Products WHERE Name LIKE '%{term}%'", conn);
        return cmd.ExecuteReader();
    }
}
