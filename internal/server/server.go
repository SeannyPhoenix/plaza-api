package server

import (
	"net/http"

	"github.com/seannyphoenix/plaza-api/internal/graphql"
)

// Create the server
func Create() *http.Server {
	// get the graphql handler
	handler, err := graphql.GQLHandler()
	if err != nil {
		panic(err)
	}
	return &http.Server{
		Addr:    ":7890",
		Handler: handler,
	}
}
