package main

import (
	"crypto/md5"
	"fmt"
)

func main() {
	var password string
	fmt.Print("Password: ")
	fmt.Scan(&password)

	// CWE-327: MD5 is cryptographically broken.
	h := md5.Sum([]byte(password))
	fmt.Printf("MD5: %x\n", h)
}
