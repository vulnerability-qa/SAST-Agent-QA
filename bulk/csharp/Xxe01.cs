// CWE-611: XXE via XmlDocument without disabling DTD
using System.Xml;
public class Xxe01 {
    public XmlDocument Parse(string xml) {
        var doc = new XmlDocument();
        doc.XmlResolver = null;
        doc.LoadXml(xml); // DTD processing enabled
        return doc;
    }
}
