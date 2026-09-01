package main

import (
    "context"
    "log"
    "time"

    "go.mongodb.org/mongo-driver/v2/mongo"
    "go.mongodb.org/mongo-driver/v2/mongo/options"
)

func main() {
    ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
    defer cancel()

    client, err := mongo.Connect(options.Client().ApplyURI("mongodb://localhost:27017"))
    if err != nil {
        log.Fatal(err)
    }
    defer client.Disconnect(ctx)

    sess, err := client.StartSession()
    if err != nil {
        log.Fatal(err)
    }
    defer sess.EndSession(ctx)

    err = mongo.WithSession(ctx, sess, func(sc context.Context) error {
        _, err := sess.WithTransaction(sc, func(sc context.Context) (any, error) {
            // call code paths you’re changing (e.g. insert, then commit/abort)
            return nil, nil
        })
        return err
    })
    log.Printf("withTransaction err = %v", err)
}
