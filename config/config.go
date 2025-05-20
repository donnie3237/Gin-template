// Config here
package config

import (
    "context"
    "fmt"
    "log"
    "os"
    "time"

    "github.com/joho/godotenv"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

func LoadEnv() {
    if err := godotenv.Load(); err != nil {
        log.Println("No .env file found, using system envs")
    }
}

func ConnectMongoDB() *mongo.Database {
    uri := os.Getenv("MONGO_URI")
    dbName := os.Getenv("MONGO_DB")

    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
    if err != nil {
        log.Fatal(err)
    }

    return client.Database(dbName)
}
