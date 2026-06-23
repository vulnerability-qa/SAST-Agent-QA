// CWE-91: XML Injection via string concatenation
public class XmlInjection01 {
    public String buildXml(String username) {
        return "<user><name>" + username + "</name></user>"; // inject extra tags
    }
}
