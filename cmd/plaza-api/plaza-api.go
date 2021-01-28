package main

import (
	"log"
	"net/http"

	"github.com/seannyphoenix/plaza-api/interal/server"
)

func main() {
	log.Fatal(http.ListenAndServe(":7890", server.App()))
}
