// CWE-22: Path Traversal via filepath.Join without validation
package main
import (
	"net/http"
	"os"
	"path/filepath"
)
func serveFile(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	path := filepath.Join("/uploads", name)
	data, _ := os.ReadFile(path)
	w.Write(data)
}
