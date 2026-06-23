// CWE-94: SSTI via FreeMarker
import freemarker.template.*;
import java.io.*;
public class SstiFreemarker01 {
    public String render(String templateStr, java.util.Map<String, Object> model) throws Exception {
        Configuration cfg = new Configuration(Configuration.VERSION_2_3_31);
        Template t = new Template("inline", new StringReader(templateStr), cfg);
        StringWriter sw = new StringWriter();
        t.process(model, sw);
        return sw.toString();
    }
}
