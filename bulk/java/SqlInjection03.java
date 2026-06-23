// CWE-89: SQL Injection via HQL/JPQL
import org.hibernate.Session;
public class SqlInjection03 {
    public Object findUser(Session session, String role) {
        return session.createQuery("FROM User WHERE role = '" + role + "'").uniqueResult();
    }
}
