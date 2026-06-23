// CWE-611: XXE via XMLInputFactory
import javax.xml.stream.*;
import java.io.*;
public class Xxe03 {
    public XMLStreamReader parse(String xml) throws Exception {
        XMLInputFactory factory = XMLInputFactory.newInstance();
        // IS_SUPPORTING_EXTERNAL_ENTITIES not disabled
        return factory.createXMLStreamReader(new StringReader(xml));
    }
}
