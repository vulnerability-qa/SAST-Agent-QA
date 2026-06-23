// CWE-90: LDAP Injection
import javax.naming.*;
import javax.naming.directory.*;
public class LdapInjection01 {
    public NamingEnumeration<?> findUser(DirContext ctx, String username) throws NamingException {
        String filter = "(uid=" + username + ")";
        return ctx.search("dc=corp,dc=com", filter, new SearchControls());
    }
}
