package graphql

import (
	"log"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
)

var allTypes = []graphql.Type{}
var allQueries = graphql.Fields{}
var allMutations = graphql.Fields{}

// GQLHandler is the base graphql handler for the go server
func GQLHandler() (*handler.Handler, error) {
	schema, err := graphql.NewSchema(graphql.SchemaConfig{
		Types: allTypes,
		Query: graphql.NewObject(
			graphql.ObjectConfig{
				Name:   "Query",
				Fields: allQueries,
			},
		),
		Mutation: graphql.NewObject(
			graphql.ObjectConfig{
				Name:   "Mutation",
				Fields: allMutations,
			},
		),
	})

	if err != nil {
		log.Fatalf("failed to create new schema, error: %v", err)
	}

	h := handler.New(&handler.Config{
		Schema: &schema,
		Pretty: true,
	})

	return h, nil
}
