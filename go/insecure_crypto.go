package main

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"golang.org/x/crypto/argon2"
)

func main() {
	var password string
	fmt.Print("Password: ")
	fmt.Scan(&password)

	memory := uint32(64 * 1024)
	iterations := uint32(3)
	parallelism := uint8(2)
	saltLength := uint32(16)
	keyLength := uint32(32)

	salt := make([]byte, saltLength)
	if _, err := rand.Read(salt); err != nil {
		fmt.Printf("Error generating salt: %v\n", err)
		return
	}

	// CWE-327 Fix: Use Argon2id for password hashing
	derivedKey := argon2.IDKey([]byte(password), salt, iterations, memory, parallelism, keyLength)
	encodedSalt := base64.RawStdEncoding.EncodeToString(salt)
	encodedKey := base64.RawStdEncoding.EncodeToString(derivedKey)
	fmt.Printf("Argon2id: $argon2id$v=%d$m=%d,t=%d,p=%d$%s$%s\n", argon2.Version, memory, iterations, parallelism, encodedSalt, encodedKey)
}
