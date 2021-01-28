package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func setContentType(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		h.ServeHTTP(w, r)
	})
}

func defaultHandle(w http.ResponseWriter, r *http.Request) {
	names := []string{"Seanny", "Brandon", "Home"}
	json.NewEncoder(w).Encode(names)
}

func main() {
	mux := http.NewServeMux()

	final := http.HandlerFunc(defaultHandle)
	mux.Handle("/", setContentType(final))
	log.Fatal(http.ListenAndServe(":7890", mux))
}
