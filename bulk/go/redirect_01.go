// CWE-601: Open Redirect
package main
import "net/http"
func loginHandler(w http.ResponseWriter, r *http.Request) {
	next := r.URL.Query().Get("next")
	if next == "" {
		next = "/"
	}
	http.Redirect(w, r, next, http.StatusFound)
}
