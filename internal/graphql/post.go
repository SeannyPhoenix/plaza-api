package graphql

import (
	"github.com/graphql-go/graphql"
)

var postType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Post",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.String,
			},
			"name": &graphql.Field{
				Type: graphql.String,
			},
			"timestamp": &graphql.Field{
				Type: graphql.String,
			},
			"contents": &graphql.Field{
				Type: graphql.String,
			},
			"account": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)

var postQueryFields = &graphql.Field{
	Type: postType,
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
	},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		return nil, nil
	},
}

func init() {
	allTypes = append(allTypes, postType)
	allQueries["post"] = postQueryFields
}
