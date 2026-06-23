// CWE-89: SQL Injection in JPA native query
import javax.persistence.*;
public class SqlInjection02 {
    @PersistenceContext EntityManager em;
    public Object findByEmail(String email) {
        return em.createNativeQuery("SELECT * FROM users WHERE email = '" + email + "'").getSingleResult();
    }
}
