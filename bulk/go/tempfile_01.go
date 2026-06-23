// CWE-377: Insecure temp file with predictable name
package main
import (
	"fmt"
	"os"
)
func writeTmp(data []byte) (*os.File, error) {
	path := fmt.Sprintf("/tmp/upload_%d", os.Getpid())
	return os.OpenFile(path, os.O_CREATE|os.O_WRONLY, 0600)
}
