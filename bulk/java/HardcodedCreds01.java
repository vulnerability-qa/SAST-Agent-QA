// CWE-798: Hardcoded database credentials
import java.sql.*;
public class HardcodedCreds01 {
    private static final String DB_PASS = "Sup3rS3cr3t!";
    public Connection getConn() throws SQLException {
        return DriverManager.getConnection("jdbc:mysql://db:3306/app", "root", DB_PASS);
    }
}
