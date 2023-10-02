package db

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/event"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func DBinstance() *mongo.Client {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	/*
		// ctx := context.Background()
		clientOpts := options.Client().ApplyURI(os.Getenv("MONGODB_URI")).SetMonitor(cmdMonitor)
		log.Println(clientOpts)
	*/

	cmdMonitor := &event.CommandMonitor{
		Started: func(_ context.Context, evt *event.CommandStartedEvent) {
			log.Print(evt.Command)
		},
	}
	mongoDbUri := os.Getenv("MONGODB_URI")
	log.Println("Connecting to MongoDB at " + mongoDbUri + " ...")

	client, err := mongo.NewClient(options.Client().ApplyURI(mongoDbUri).SetMonitor(cmdMonitor))
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Connected to MongoDB !")
	return client
}

var Client *mongo.Client = DBinstance()

func OpenCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	var collection *mongo.Collection = client.Database(os.Getenv("MONGO_DB_NAME")).Collection(collectionName)
	return collection
}
