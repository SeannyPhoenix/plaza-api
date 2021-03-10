package db

import (
	"errors"

	"github.com/google/uuid"
	"github.com/graphql-go/graphql"
	"github.com/seannyphoenix/plaza-api/internal/plazaerr"
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
	// accept any one of these keys to search preferred in this order:
	// id, email, name
	filter := bson.E{}
	if val, ok := p.Args["id"]; ok {
		filter.Key = "_id"
		filter.Value = val.(string)
	} else if val, ok := p.Args["email"]; ok {
		filter.Key = "email"
		filter.Value = val.(string)
	} else if val, ok := p.Args["name"]; ok {
		filter.Key = "name"
		filter.Value = val.(string)
	}

	var account Account
	err := accountCollection.FindOne(p.Context, bson.D{filter}).Decode(&account)

	if err == mongo.ErrNoDocuments {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return account, nil
}

// CreateAccount ... create an account
func CreateAccount(p graphql.ResolveParams) (interface{}, error) {
	accountDetails := Account{
		ID: uuid.NewString(),
	}

	if val, ok := p.Args["name"]; ok {
		accountDetails.Name = val.(string)
	} else {
		return nil, errors.New("must include name argument")
	}
	if val, ok := p.Args["email"]; ok {
		accountDetails.Email = val.(string)
	} else {
		return nil, errors.New("must include email argument")
	}

	// Before we try to create a new account, check to make sure
	// there is no account using this name or email
	query := bson.D{{Key: "$or", Value: []bson.D{
		{{Key: "name", Value: accountDetails.Name}},
		{{Key: "email", Value: accountDetails.Email}},
	}}}

	var existing Account
	err := accountCollection.FindOne(p.Context, query).Decode(&existing)

	// var cur *mongo.Cursor
	// if cursor, err := accountCollection.Find(p.Context, query); err != nil {
	// 	fmt.Printf("%+v\n", err)
	// } else {
	// 	cur = cursor
	// }

	// Only create an acount if there isn't one already
	if err == mongo.ErrNoDocuments {
		_, err := accountCollection.InsertOne(p.Context, accountDetails)
		if err != nil {
			return nil, err
		}
		return accountDetails, nil
	}

	// Otherwise, return nothing
	return nil, plazaerr.DuplicateAccount("name", "Seanny")
}

// UpdateAccount ... updates an account by ID
func UpdateAccount(p graphql.ResolveParams) (interface{}, error) {
	// var id string
	// filter := bson.E{}

	// var accountDetails Account
	// for name, value := range p.Args {
	// 	accountDetails = value
	// 	fmt.Printf("%+v: %+v\n", name, value)
	// }

	// if val, ok := p.Args["id"]; ok {
	// 	filter.Key = "_id"
	// 	filter.Value = val.(string)
	// 	id = val.(string)
	// } else {
	// 	return false, nil
	// }

	// var account Account
	// err := accountCollection.FindOneAndUpdate(p.Context, bson.D{filter}).Decode(&account)
	// if err != mongo.ErrNoDocuments && account.ID == id {
	// 	return true, nil
	// }
	return nil, nil
}

// DeleteAccount ... deletes an account by ID
func DeleteAccount(p graphql.ResolveParams) (interface{}, error) {
	var id string
	filter := bson.E{}

	if val, ok := p.Args["id"]; ok {
		filter.Key = "_id"
		filter.Value = val.(string)
		id = val.(string)
	} else {
		return false, nil
	}

	var account Account
	err := accountCollection.FindOneAndDelete(p.Context, bson.D{filter}).Decode(&account)
	if err != mongo.ErrNoDocuments && account.ID == id {
		return true, nil
	}
	return false, err
}
