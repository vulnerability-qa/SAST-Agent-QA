// CWE-78: Command Injection via ProcessBuilder with shell
import java.io.*;
import java.util.*;
public class CmdInjection02 {
    public void convert(String file) throws IOException {
        new ProcessBuilder("sh", "-c", "convert " + file + " output.png").start();
    }
}
