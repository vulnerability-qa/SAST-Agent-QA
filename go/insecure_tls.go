package main

import (
	"crypto/tls"
	"fmt"
	"net/http"
)

func main() {
	// CWE-295: TLS certificate verification disabled — vulnerable to MITM attacks.
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	resp, err := client.Get("https://example.com")
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Println(resp.Status)
}
