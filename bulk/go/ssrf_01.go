// CWE-918: SSRF via http.Get with user-controlled URL
package main
import "net/http"
func fetchURL(url string) (*http.Response, error) {
	return http.Get(url) // no allowlist validation
}
