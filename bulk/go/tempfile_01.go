// CWE-377: Insecure temp file with predictable name
package main
import (
	"fmt"
	"os"
)
func writeTmp(data []byte) (*os.File, error) {
	f, err := os.CreateTemp("/tmp", "upload_*")
	if err != nil {
		return nil, err
	}
	return f, nil
}
