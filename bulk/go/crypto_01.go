// CWE-327: MD5 used for sensitive hashing
package main
import (
	"crypto/md5"
	"fmt"
)
func hashPassword(password string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(password)))
}
