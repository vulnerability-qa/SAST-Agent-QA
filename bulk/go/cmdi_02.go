// CWE-78: Command Injection via os.system equivalent
package main
import "os/exec"
func convert(filename string) error {
	cmd := exec.Command("bash", "-c", "convert "+filename+" output.png")
	return cmd.Run()
}
