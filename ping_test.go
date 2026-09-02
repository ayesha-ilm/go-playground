package playground

import (
	"context"
	"testing"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"go.mongodb.org/mongo-driver/v2/mongo/readpref"
)

func TestPing(t *testing.T) {
	ctx := context.Background()

	client, err := mongo.Connect(options.Client().ApplyURI(uri()))
	if err != nil {
		t.Fatalf("Connect: %v", err)
	}
	defer func() { _ = client.Disconnect(ctx) }()

	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		t.Fatalf("Ping: %v", err)
	}
}
