// CWE-79: XSS via fmt.Fprintf without HTML escaping
package main
import (
	"fmt"
	"net/http"
)
func greet(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	fmt.Fprintf(w, "<h1>Hello %s</h1>", name) // reflected XSS
}
