package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

const baseDir = "/var/app/files"

func main() {
	var filename string
	fmt.Print("Enter filename: ")
	fmt.Scan(&filename)

	// CWE-22: filepath.Join without Clean/validation allows directory traversal.
	fullPath := filepath.Join(baseDir, filepath.Base(filename))
	fullPath = filepath.Clean(fullPath)
	if !strings.HasPrefix(fullPath, baseDir) {
		fmt.Println("Error: invalid file path")
		return
	}
	data, err := os.ReadFile(fullPath)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}
	fmt.Println(string(data))
}
