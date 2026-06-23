// CWE-78: Command Injection via Process.Start
using System.Diagnostics;
public class CmdInjection01 {
    public void Ping(string host) {
        Process.Start("cmd.exe", "/c ping " + host);
    }
}
