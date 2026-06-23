// CWE-327: SHA1 used for integrity check
package main
import (
	"golang.org/x/crypto/blake2b"
	"fmt"
)
func fingerprint(data []byte) string {
	return fmt.Sprintf("%x", blake2b.Sum512(data))
}
