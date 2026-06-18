import java.sql.Connection;
import java.sql.DriverManager;
import java.sql.ResultSet;
import java.sql.Statement;
import java.util.Scanner;

public class SqlInjection {
    public static void main(String[] args) throws Exception {
        Connection conn = DriverManager.getConnection("jdbc:h2:mem:test");
        conn.createStatement().execute("CREATE TABLE users (name VARCHAR, password VARCHAR)");
        conn.createStatement().execute("INSERT INTO users VALUES ('alice','wonderland')");

        Scanner sc = new Scanner(System.in);
        System.out.print("User: ");
        String username = sc.nextLine();

        // CWE-89: user input concatenated directly into SQL query.
        Statement stmt = conn.createStatement();
        ResultSet rs = stmt.executeQuery("SELECT * FROM users WHERE name = '" + username + "'");
        while (rs.next()) {
            System.out.println(rs.getString(1));
        }
    }
}
