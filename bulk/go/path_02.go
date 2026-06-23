// CWE-22: Path Traversal in file write
package main
import "os"
func saveUpload(filename string, data []byte) error {
	return os.WriteFile("/var/www/uploads/"+filename, data, 0644)
}
