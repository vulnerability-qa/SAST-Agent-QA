package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/greet", func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Query().Get("name")
		w.Header().Set("Content-Type", "text/html")
		// CWE-79: user input written to response without HTML encoding.
		fmt.Fprintf(w, "<h1>Hello, %s!</h1>", name)
	})
	srv := &http.Server{
		Addr:              ":8080",
		ReadHeaderTimeout: 15 * time.Second,
		ReadTimeout:       15 * time.Second,
		WriteTimeout:      10 * time.Second,
		IdleTimeout:       30 * time.Second,
	}
	srv.ListenAndServe()
}
