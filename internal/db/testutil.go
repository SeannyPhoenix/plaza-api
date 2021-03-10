package db

import (
	"context"
	"testing"
)

// Initialize before tests
func initialize(t *testing.T) {
	// t.Parallel()
	t.Cleanup(cleanup)
	if !IsConnected() {
		Connect(context.Background(), "mongodb://mongo.seannyphoenix.com:27017", "plaza")
	}
}

// Cleanup after tests
func cleanup() {
	if IsConnected() {
		Disconnect(context.Background())
	}
}
