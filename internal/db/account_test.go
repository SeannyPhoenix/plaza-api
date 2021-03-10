package db

import (
	"context"
	"testing"

	"github.com/google/uuid"
)

var a Account
var ctx context.Context

func TestAccount(t *testing.T) {
	initialize(t)
	t.Run("Create", createAccount)
	t.Run("Get", getAccount)
}

func createAccount(t *testing.T) {
	if err := CreateAccount(ctx, a); err != nil {
		t.Fatal(err)
	}
	if _, err := GetAccountByID(ctx, a.ID); err != nil {
		t.Fatal(err)
	}
}

func getAccount(t *testing.T) {
	if _, err := GetAccountByID(ctx, a.ID); err != nil {
		t.Fatal(err)
	}
}

func init() {
	a = Account{
		ID:    uuid.New(),
		Name:  "Seanny Phoenix",
		Email: "seannyphoenix@gmail.com",
	}
	ctx = context.Background()
}
