// CWE-113: HTTP Response Splitting via Header injection
package main
import "net/http"
func setLang(w http.ResponseWriter, r *http.Request) {
	lang := r.URL.Query().Get("lang")
	w.Header().Set("X-Language", lang) // CRLF injection
}
