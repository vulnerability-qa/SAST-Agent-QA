// CWE-502: Insecure BinaryFormatter deserialization
using System.IO;
using System.Runtime.Serialization.Formatters.Binary;
public class Deserialization01 {
    public object Deserialize(byte[] data) {
        var formatter = new BinaryFormatter();
        return formatter.Deserialize(new MemoryStream(data));
    }
}
