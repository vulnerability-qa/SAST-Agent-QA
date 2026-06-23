// CWE-22: Path Traversal in file write
package main
import (
	"os"
	"path/filepath"
)
func saveUpload(filename string, data []byte) error {
	sanitized := filepath.Base(filename)
	fullPath := filepath.Join("/var/www/uploads", sanitized)
	return os.WriteFile(fullPath, data, 0600)
}
