// CWE-90: LDAP Injection
using System.DirectoryServices;
public class LdapInjection01 {
    public SearchResultCollection FindUser(string username) {
        var entry = new DirectoryEntry("LDAP://dc=corp,dc=com");
        var searcher = new DirectorySearcher(entry);
        searcher.Filter = "(uid=" + username + ")"; // injection possible
        return searcher.FindAll();
    }
}
