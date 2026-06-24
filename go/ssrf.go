package main

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net"
	"net/http"
	"time"
)

func IsDisallowedIP(hostIP string) bool {
	ip := net.ParseIP(hostIP)
	return ip.IsMulticast() || ip.IsUnspecified() || ip.IsLoopback() || ip.IsPrivate()
}

func SafeTransport(timeout time.Duration) *http.Transport {
	return &http.Transport{
		DialContext: func(ctx context.Context, network, addr string) (net.Conn, error) {
			c, err := net.DialTimeout(network, addr, timeout)
			if err != nil {
				return nil, err
			}
			ip, _, _ := net.SplitHostPort(c.RemoteAddr().String())
			if IsDisallowedIP(ip) {
				c.Close()
				return nil, errors.New("ip address is not allowed")
			}
			return c, err
		},
	}
}

func main() {
	http.HandleFunc("/fetch", func(w http.ResponseWriter, r *http.Request) {
		url := r.URL.Query().Get("url")
		const clientConnectTimeout = 10 * time.Second
		httpClient := &http.Client{Transport: SafeTransport(clientConnectTimeout)}
		// CWE-918: user-controlled URL fetched without validation allows SSRF.
		resp, err := httpClient.Get(url)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		defer resp.Body.Close()
		io.Copy(w, resp.Body)
	})
	srv := &http.Server{
		Addr:              ":8080",
		ReadHeaderTimeout: 15 * time.Second,
		ReadTimeout:       15 * time.Second,
		WriteTimeout:      10 * time.Second,
		IdleTimeout:       30 * time.Second,
	}
	fmt.Println(srv.ListenAndServe())
}
