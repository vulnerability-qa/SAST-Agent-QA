// CWE-918: SSRF via http.Client with user-supplied URL
package main
import "net/http"
func proxy(w http.ResponseWriter, r *http.Request) {
	target := r.URL.Query().Get("url")
	resp, _ := http.Get(target)
	if resp != nil {
		defer resp.Body.Close()
	}
}
