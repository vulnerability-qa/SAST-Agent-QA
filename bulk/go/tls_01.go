// CWE-295: TLS certificate verification disabled
package main
import (
	"crypto/tls"
	"net/http"
)
func insecureClient() *http.Client {
	return &http.Client{}
}
