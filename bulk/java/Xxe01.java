// CWE-611: XXE via DocumentBuilder
import javax.xml.parsers.*;
import org.w3c.dom.*;
import java.io.*;
public class Xxe01 {
    public Document parse(InputStream is) throws Exception {
        DocumentBuilder builder = DocumentBuilderFactory.newInstance().newDocumentBuilder();
        return builder.parse(is); // XXE enabled by default
    }
}
