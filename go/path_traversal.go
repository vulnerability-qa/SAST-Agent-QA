package main

import (
	"fmt"
	"os"
	"path/filepath"
)

const baseDir = "/var/app/files"

func main() {
	var filename string
	fmt.Print("Enter filename: ")
	fmt.Scan(&filename)

	// CWE-22: filepath.Join without Clean/validation allows directory traversal.
	fullPath := filepath.Join(baseDir, filename)
	data, _ := os.ReadFile(fullPath)
	fmt.Println(string(data))
}
