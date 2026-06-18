package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/fetch", func(w http.ResponseWriter, r *http.Request) {
		url := r.URL.Query().Get("url")
		// CWE-918: user-controlled URL fetched without validation allows SSRF.
		resp, err := http.Get(url)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		defer resp.Body.Close()
		io.Copy(w, resp.Body)
	})
	fmt.Println(http.ListenAndServe(":8080", nil))
}
