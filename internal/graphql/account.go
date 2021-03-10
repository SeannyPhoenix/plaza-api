package graphql

import (
	"github.com/graphql-go/graphql"
	"github.com/seannyphoenix/plaza-api/internal/db"
)

// GraphQL Account type
var accountType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Account",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.String,
			},
			"name": &graphql.Field{
				Type: graphql.String,
			},
			"email": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)

var getAccount = &graphql.Field{
	Type: accountType,
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"name": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"email": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
	},
	Resolve: db.GetAccount,
}

var createAccount = &graphql.Field{
	Type: accountType,
	Args: graphql.FieldConfigArgument{
		"name": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"email": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
	},
	Resolve: db.CreateAccount,
}

var updateAccount = &graphql.Field{
	Type: graphql.String,
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"name": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"email": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
	},
	Resolve: db.UpdateAccount,
}

var deleteAccount = &graphql.Field{
	Type: graphql.Boolean,
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
	},
	Resolve: db.DeleteAccount,
}

func init() {
	allTypes = append(allTypes, accountType)
	// allQueries["me"] = me
	allQueries["getAccount"] = getAccount
	allMutations["createAccount"] = createAccount
	allMutations["updateAccount"] = updateAccount
	allMutations["deleteAccount"] = deleteAccount
}
