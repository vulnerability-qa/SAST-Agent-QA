// CWE-377: Insecure Temporary File creation
import java.io.*;
public class TempFile01 {
    public File createTemp(String prefix) throws IOException {
        File f = new File("/tmp/" + prefix + System.currentTimeMillis());
        f.createNewFile();
        return f;
    }
}
