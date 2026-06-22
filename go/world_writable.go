package main

import (
	"fmt"
	"os"
)

func main() {
	// CWE-276: file created with world-writable permissions (0777).
	f, err := os.CreateTemp("/tmp", "output-*.txt")
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	defer f.Close()
	
	_, err = f.Write([]byte("sensitive data"))
	if err != nil {
		fmt.Println("error:", err)
		return
	}
}
