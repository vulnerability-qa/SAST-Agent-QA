// CWE-327: SHA1 used for integrity check
package main
import (
	"crypto/sha1"
	"fmt"
)
func fingerprint(data []byte) string {
	return fmt.Sprintf("%x", sha1.Sum(data))
}
