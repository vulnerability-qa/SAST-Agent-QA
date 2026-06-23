// CWE-639: IDOR in HTTP handler
package main
import (
	"net/http"
	"fmt"
)
func getInvoice(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	// No ownership check — any authenticated user can access any invoice
	fmt.Fprintf(w, "invoice data for %s", id)
}
