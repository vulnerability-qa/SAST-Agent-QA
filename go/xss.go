package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/greet", func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Query().Get("name")
		w.Header().Set("Content-Type", "text/html")
		// CWE-79: user input written to response without HTML encoding.
		fmt.Fprintf(w, "<h1>Hello, %s!</h1>", name)
	})
	http.ListenAndServe(":8080", nil)
}
