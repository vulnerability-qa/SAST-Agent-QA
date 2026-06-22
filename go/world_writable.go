package main

import (
	"fmt"
	"os"
)

func main() {
	// CWE-276: file created with world-writable permissions (0777).
	err := os.WriteFile("/tmp/output.txt", []byte("sensitive data"), 0600)
	if err != nil {
		fmt.Println("error:", err)
	}
}
