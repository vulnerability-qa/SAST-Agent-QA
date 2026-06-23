// CWE-526: Sensitive data in environment variable exposed via HTTP
package main
import (
	"net/http"
	"os"
)
func debugHandler(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Query().Get("key")
	w.Write([]byte(os.Getenv(key))) // exposes arbitrary env vars
}
