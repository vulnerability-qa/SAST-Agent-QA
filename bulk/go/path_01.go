// CWE-22: Path Traversal via filepath.Join without validation
package main
import (
	"net/http"
	"os"
	"path/filepath"
	"strings"
)
func serveFile(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	safeName := filepath.Base(name)
	path := filepath.Join("/uploads", safeName)

	baseDir, err := filepath.Abs("/uploads")
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	absPath, err := filepath.Abs(path)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	if !strings.HasPrefix(absPath, baseDir + string(filepath.Separator)) {
		http.Error(w, "Invalid file path", http.StatusBadRequest)
		return
	}

	data, err := os.ReadFile(absPath)
	if err != nil {
		http.Error(w, "File not found", http.StatusNotFound)
		return
	}
	w.Write(data)
}
