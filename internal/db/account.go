package db

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/graphql-go/graphql"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// Account ... Account Model for DB
type Account struct {
	ID    string `bson:"_id" json:"id"`
	Name  string `bson:"name" json:"name"`
	Email string `bson:"email" json:"email"`
}

// GetAccount ... gets an account by id, name, or email
func GetAccount(p graphql.ResolveParams) (interface{}, error) {
	accountDetails := Account{
		ID:    p.Args["id"].(string),
		Name:  p.Args["name"].(string),
		Email: p.Args["email"].(string),
	}
	fmt.Printf("%+v", accountDetails)
	// return accountDetails, nil
	return accountDetails, nil
}

// CreateAccount ... create an account
func CreateAccount(p graphql.ResolveParams) (interface{}, error) {
	accountDetails := Account{
		ID:    uuid.NewString(),
		Name:  p.Args["name"].(string),
		Email: p.Args["email"].(string),
	}

	var existing Account

	query := bson.D{{Key: "$or", Value: []bson.D{
		{{Key: "name", Value: accountDetails.Name}},
		{{Key: "email", Value: accountDetails.Email}},
	}}}

	err := accountCollection.FindOne(p.Context, query).Decode(&existing)

	// Only create an acount if there isn't one already
	if err == mongo.ErrNoDocuments {
		newAccount, err := accountCollection.InsertOne(p.Context, accountDetails)
		if err != nil {
			fmt.Printf("%+v\n", err)
			return nil, err
		}
		return newAccount.InsertedID, nil
	}

	// Otherwise, return nothing
	fmt.Printf("%+v\n", err)
	return nil, err
}
