// CWE-798: Hardcoded connection string
using System.Data.SqlClient;
public class HardcodedCreds01 {
    private const string ConnStr = "Server=db;Database=app;User=sa;Password=P@ssw0rd!";
    public SqlConnection GetConnection() => new SqlConnection(ConnStr);
}
