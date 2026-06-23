// CWE-502: Insecure Deserialization via ObjectInputStream
import java.io.*;
public class Deserialization01 {
    public Object deserialize(byte[] data) throws Exception {
        ObjectInputStream ois = new ObjectInputStream(new ByteArrayInputStream(data));
        return ois.readObject(); // arbitrary class instantiation
    }
}
