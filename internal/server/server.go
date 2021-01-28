package server

import (
	"fmt"
	"net/http"
)

// App is the main server
func App() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, "Terrible news")
	})
}
