package db

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/seannyphoenix/plazaapi/internal/plazaerr"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Account ... x
type Account struct {
	ID     uuid.UUID `bson:"_id" json:"id"`
	Name   string    `bson:"name" json:"name"`
	Email  string    `bson:"email" json:"email"`
	Status string    `bson:"status" json:"status"`
}

type accountUpdate struct {
	ID     uuid.UUID `bson:"-"`
	Name   string    `bson:"name,omitempty"`
	Email  string    `bson:"email,omitempty"`
	Status string    `bson:"status,omitempty"`
}

// InitializeAccountCollection ... x
func InitializeAccountCollection(ctx context.Context) error {
	indexMod := mongo.IndexModel{
		Keys: bson.D{
			{Key: "name", Value: 1},
			{Key: "email", Value: 1},
		},
		Options: options.Index().SetUnique(true),
	}
	_, err := collection("account").Indexes().CreateOne(ctx, indexMod)
	return err
}

// CreateAccount ... x
func CreateAccount(ctx context.Context, a Account) error {
	_, err := collection("account").InsertOne(ctx, a)
	if mongo.IsDuplicateKeyError(err) {
		return err
	}
	return nil
}

// GetAccountByID ... x
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
	update := bson.M{"$set": bson.M{"status": "inactive"}}
	result := collection("account").FindOneAndUpdate(ctx, bson.M{"_id": acctID}, update)
	return result.Err()
}

// UpdateAccount ... x
func UpdateAccount(ctx context.Context, a Account) error {
	filter := bson.M{"_id": a.ID}
	update := bson.M{"$set": accountUpdate(a)}
	var acct Account
	err := collection("account").FindOneAndUpdate(
		ctx,
		filter,
		update,
	).Decode(&acct)
	fmt.Printf("%+v\n", err)
	if err == mongo.ErrNoDocuments {
		return plazaerr.ErrNotFound.Errorf("no account found for id: %s", a.ID.String())
	}
	return err
}
