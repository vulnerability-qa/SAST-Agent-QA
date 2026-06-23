// CWE-295: TLS certificate verification disabled
package main
import (
	"crypto/tls"
	"net/http"
)
func insecureClient() *http.Client {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	return &http.Client{Transport: tr}
}
