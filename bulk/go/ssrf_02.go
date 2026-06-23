// CWE-918: SSRF via http.Client with user-supplied URL
package main
import (
	"context"
	"errors"
	"net"
	"net/http"
	"time"
)

// IsDisallowedIP parses the ip to determine if we should allow the HTTP client to continue
func IsDisallowedIP(hostIP string) bool {
	ip := net.ParseIP(hostIP)
	return ip.IsMulticast() || ip.IsUnspecified() || ip.IsLoopback() || ip.IsPrivate()
}

// SafeTransport uses the net.Dial to connect, then if successful check if the resolved
// ip address is disallowed. We do this due to hosts such as localhost.lol being resolvable to
// potentially malicious URLs. We allow connection only for resolution purposes.
func SafeTransport(timeout time.Duration) *http.Transport {
	return &http.Transport{
		DialContext: func(ctx context.Context, network, addr string) (net.Conn, error) {
			c, err := net.DialTimeout(network, addr, timeout)
			if err != nil {
				return nil, err
			}
			ip, _, _ := net.SplitHostPort(c.RemoteAddr().String())
			if IsDisallowedIP(ip) {
				return nil, errors.New("ip address is not allowed")
			}
			return c, err
		},
		TLSHandshakeTimeout: timeout,
	}
}

func proxy(w http.ResponseWriter, r *http.Request) {
	target := r.URL.Query().Get("url")
	const clientConnectTimeout = time.Second * 10
	httpClient := &http.Client{
		Transport: SafeTransport(clientConnectTimeout),
	}
	resp, _ := httpClient.Get(target)
	if resp != nil {
		defer resp.Body.Close()
	}
}
