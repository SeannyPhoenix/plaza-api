package db

import (
	"context"

	"github.com/google/uuid"
	"github.com/seannyphoenix/golang-graphql-mongodb/internal/plazaerr"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// Account ... xx
type Account struct {
	ID     uuid.UUID `bson:"_id" json:"id"`
	Name   string    `bson:"name" json:"name"`
	Email  string    `bson:"email" json:"email"`
	Status string    `bson:"status" json:"status"`
}

// CreateAccount ... x
func CreateAccount(ctx context.Context, a Account) error {
	_, err := collection("account").InsertOne(ctx, a)
	if dup, index := checkDup(err); dup {
		switch index {
		case "name":
			return plazaerr.DuplicateAccount(index, "")
		}
	}
	return err
}

// GetAccountByID ... xxx
func GetAccountByID(ctx context.Context, acctID uuid.UUID) (acct Account, err error) {
	filter := bson.M{"_id": acctID}
	err = collection("account").FindOne(
		ctx,
		filter,
	).Decode(&acct)
	if err == mongo.ErrNoDocuments {
		err = plazaerr.ErrNotFound.Errorf("no account found for id: %s", acctID.String())
	}
	return
}

// DeleteAccount ... x
func DeleteAccount(ctx context.Context, acctID uuid.UUID) error {
	update := bson.M{"$set": bson.M{"status": false}}

	result := collection("account").FindOneAndUpdate(ctx, bson.M{"_id": acctID}, update)
	return result.Err()
}

// UpdateAccount ... x
func UpdateAccount(ctx context.Context, a Account) error {
	filter := bson.M{"_id": a.ID}
	// update := bson.M{}
	var acct Account
	err := collection("account").FindOneAndUpdate(
		ctx,
		filter,
		nil,
	).Decode(acct)
	if err == mongo.ErrNoDocuments {
		err = plazaerr.ErrNotFound.Errorf("no account found for id: %s", a.ID.String())
	}
	return nil
}
