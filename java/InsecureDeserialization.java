import java.io.FileInputStream;
import java.io.ObjectInputStream;

public class InsecureDeserialization {
    public static void main(String[] args) throws Exception {
        // CWE-502: deserializing untrusted data from a file allows arbitrary code execution.
        ObjectInputStream ois = new ObjectInputStream(new FileInputStream(args[0]));
        ois.setObjectInputFilter(ObjectInputFilter.Config.createFilter("java.lang.String;!*"));
        Object obj = ois.readObject();
        System.out.println(obj);
        ois.close();
    }
}
