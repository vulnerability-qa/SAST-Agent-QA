// CWE-502: Kryo deserialization without type registration
import com.esotericsoftware.kryo.Kryo;
import com.esotericsoftware.kryo.io.Input;
public class Deserialize02 {
    Kryo kryo = new Kryo();
    public Object read(byte[] data) {
        kryo.setRegistrationRequired(false); // any class can be deserialized
        return kryo.readClassAndObject(new Input(data));
    }
}
