package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter command: ")
	cmdLine, _ := reader.ReadString('\n')

	// CWE-78: user input executed via shell without validation.
	cmd := exec.Command("/bin/sh", "-c", cmdLine)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
}
