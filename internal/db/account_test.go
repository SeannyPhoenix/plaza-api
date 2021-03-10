package db

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/mongo"
)

var acctOne Account
var acctTwo Account
var ctx context.Context

func TestAccount(t *testing.T) {
	initialize(t)
	t.Run("Initialize", initAccountColl)
	t.Run("Create", createAccount)
	t.Run("Get", getAccount)
	t.Run("Update", updateAccount)
	t.Run("Delete", deleteAccount)
}

func initAccountColl(t *testing.T) {
	if err := InitializeAccountCollection(context.Background()); err != nil {
		t.Fatal(err)
	}
}

func createAccount(t *testing.T) {
	if err := CreateAccount(ctx, acctOne); err != nil {
		t.Fatal(err)
	} else if acct, err := GetAccountByID(ctx, acctOne.ID); err != nil {
		t.Fatal(err)
	} else if !isEqual(acct, acctOne) {
		t.Fatal("not equal")
	}

	if err := CreateAccount(ctx, acctOne); err != nil {
		if !mongo.IsDuplicateKeyError(err) {
			t.Fatal(err)
		}
	}
}

func getAccount(t *testing.T) {
	if acct, err := GetAccountByID(ctx, acctOne.ID); err != nil {
		t.Fatal(err)
	} else if !isEqual(acct, acctOne) {
		t.Fatal("not equal")
	}

	if _, err := GetAccountByID(ctx, uuid.New()); err == nil {
		t.Fatal("found an account?!")
	}
}

func updateAccount(t *testing.T) {
	if err := UpdateAccount(ctx, acctTwo); err != nil {
		t.Fatal(err)
	} else if acct, err := GetAccountByID(ctx, acctOne.ID); err != nil {
		t.Fatal(err)
	} else if !isEqual(acct, acctTwo) {
		t.Fatal("not equal")
	}

	if err := UpdateAccount(ctx, Account{ID: uuid.New(), Name: "test"}); err == nil {
		t.Fatal(err)
	}
}

func deleteAccount(t *testing.T) {
	if err := DeleteAccount(ctx, acctOne.ID); err != nil {
		t.Fatal(err)
	} else if acct, err := GetAccountByID(ctx, acctOne.ID); err != nil {
		t.Fatal(err)
	} else if acct.Status != "inactive" {
		t.Fatal("not inactive")
	}
}

func isEqual(a Account, b Account) bool {
	return a.ID == b.ID && a.Name == b.Name && a.Email == b.Email && a.Status == b.Status
}

func init() {
	acctOne = Account{
		ID:     uuid.New(),
		Name:   "acct_one",
		Email:  "acct_one@seannyphoenix.com",
		Status: "active",
	}

	acctTwo = Account{
		ID:     acctOne.ID,
		Name:   "acct_two",
		Email:  "acct_two@seannyphoenix.com",
		Status: "active",
	}

	ctx = context.Background()
}
