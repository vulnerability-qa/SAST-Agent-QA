// CWE-327: MD5 used for sensitive hashing
package main
import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"golang.org/x/crypto/argon2"
	"io"
)
func hashPassword(password string) (string, error) {
	// Argon2id parameters
	memory := uint32(64 * 1024)
	iterations := uint32(3)
	parallelism := uint8(2)
	saltLength := uint32(16)
	keyLength := uint32(32)

	// Generate random salt
	salt := make([]byte, saltLength)
	if _, err := io.ReadFull(rand.Reader, salt); err != nil {
		return "", err
	}

	// Create key hash derived from password
	derivedKey := argon2.IDKey([]byte(password), salt, iterations, memory, parallelism, keyLength)

	// Encode salt and key in base64
	encodedSalt := base64.RawStdEncoding.EncodeToString(salt)
	encodedKey := base64.RawStdEncoding.EncodeToString(derivedKey)

	// Return Argon2id PHC string format
	return fmt.Sprintf("$argon2id$v=%d$m=%d,t=%d,p=%d$%s$%s",
		argon2.Version,
		memory,
		iterations,
		parallelism,
		encodedSalt,
		encodedKey,
	), nil
}
