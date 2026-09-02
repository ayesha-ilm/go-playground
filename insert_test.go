package playground

import (
	"context"
	"os"
	"testing"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

// uri returns the MongoDB URI to test against, overridable with MONGODB_URI.
func uri() string {
	if u := os.Getenv("MONGODB_URI"); u != "" {
		return u
	}
	return "mongodb://localhost:27017"
}

func TestInsertMany(t *testing.T) {
	ctx := context.Background()

	client, err := mongo.Connect(options.Client().ApplyURI(uri()))
	if err != nil {
		t.Fatalf("Connect: %v", err)
	}
	defer func() { _ = client.Disconnect(ctx) }()

	coll := client.Database("playground").Collection("insert_many")
	if err := coll.Drop(ctx); err != nil {
		t.Fatalf("Drop: %v", err)
	}

	docs := []any{
		bson.D{{Key: "x", Value: 1}},
		bson.D{{Key: "x", Value: 2}},
		bson.D{{Key: "x", Value: 3}},
	}

	res, err := coll.InsertMany(ctx, docs)
	if err != nil {
		t.Fatalf("InsertMany: %v", err)
	}
	if got, want := len(res.InsertedIDs), len(docs); got != want {
		t.Fatalf("inserted %d ids, want %d", got, want)
	}

	count, err := coll.CountDocuments(ctx, bson.D{})
	if err != nil {
		t.Fatalf("CountDocuments: %v", err)
	}
	if count != int64(len(docs)) {
		t.Fatalf("count = %d, want %d", count, len(docs))
	}
}
