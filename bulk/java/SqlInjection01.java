// CWE-89: SQL Injection via Statement
import java.sql.*;
public class SqlInjection01 {
    public ResultSet getUser(Connection conn, String username) throws SQLException {
        Statement stmt = conn.createStatement();
        return stmt.executeQuery("SELECT * FROM users WHERE username = '" + username + "'");
    }
}
