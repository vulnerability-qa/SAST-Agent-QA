// CWE-470: Unsafe Reflection
public class UnsafeReflect01 {
    public Object getInstance(String className) throws Exception {
        return Class.forName(className).getDeclaredConstructor().newInstance();
    }
}
