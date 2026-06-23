// CWE-90: LDAP Injection
using System.DirectoryServices;
public class LdapInjection01 {
    public SearchResultCollection FindUser(string username) {
        string encodedUsername = EncodeLDAPString(username);
        var entry = new DirectoryEntry("LDAP://dc=corp,dc=com");
        var searcher = new DirectorySearcher(entry);
        searcher.Filter = "(uid=" + encodedUsername + ")"; // injection possible
        return searcher.FindAll();
    }

    private static string EncodeLDAPString(string input) {
        char[] chars = new char[] { '\\', '\0', '(', ')', '*' };
        string[] encoded = new string[] { "\\5c", "\\00", "\\28", "\\29", "\\2a" };

        for (int i = 0; i < chars.Length; i++)
        {
            input = input.Replace(chars[i].ToString(), encoded[i]);
        }

        return input;
    }
}
