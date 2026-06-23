// CWE-78: Command Injection via exec.Command with shell
package main
import "os/exec"
func ping(host string) ([]byte, error) {
	return exec.Command("sh", "-c", "ping -c 1 "+host).Output()
}
