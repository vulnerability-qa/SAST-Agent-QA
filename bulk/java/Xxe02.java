// CWE-611: XXE via SAXParser
import javax.xml.parsers.*;
import java.io.*;
public class Xxe02 {
    public void parse(String xml) throws Exception {
        SAXParserFactory.newInstance().newSAXParser().parse(
            new ByteArrayInputStream(xml.getBytes()), null
        );
    }
}
