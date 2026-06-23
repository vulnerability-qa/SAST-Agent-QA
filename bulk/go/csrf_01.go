// CWE-352: CSRF — state-changing endpoint with no token check
package main
import (
	"encoding/json"
	"net/http"
)
func transferHandler(w http.ResponseWriter, r *http.Request) {
	var req struct{ To string; Amount float64 }
	json.NewDecoder(r.Body).Decode(&req)
	// No CSRF token validation
	transfer(req.To, req.Amount)
	w.WriteHeader(http.StatusOK)
}
func transfer(string, float64) {}
