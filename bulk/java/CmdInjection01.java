// CWE-78: Command Injection via Runtime.exec
import java.io.*;
public class CmdInjection01 {
    public void ping(String host) throws IOException {
        Runtime.getRuntime().exec("ping -c 1 " + host);
    }
}
