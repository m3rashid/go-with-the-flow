package flow

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func StartWatchMongo(collectionNames []string) {
	client, err := mongo.NewClient(options.Client().ApplyURI(os.Getenv("MONGODB_URI")))
	if err != nil {
		panic(err)
	}

	err = client.Connect(context.Background())
	if err != nil {
		panic(err)
	}
	defer client.Disconnect(context.Background())

	for _, collectionName := range collectionNames {
		collection := client.Database(os.Getenv("MONGO_DB_NAME")).Collection(collectionName)
		go watchCollection(collection, bson.D{
			primitive.E{Key: "$match", Value: bson.D{
				primitive.E{Key: "operationType", Value: bson.D{
					primitive.E{Key: "$in", Value: bson.A{"insert", "update", "delete"}},
				}},
			}},
		})
	}

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, os.Interrupt, syscall.SIGTERM)
	<-sigs
	log.Println("Exiting...")
}

func watchCollection(collection *mongo.Collection, changeStreamStage bson.D) {
	pipeline := mongo.Pipeline{changeStreamStage}
	changeStream, err := collection.Watch(context.Background(), pipeline, nil)
	if err != nil {
		panic(err)
	}
	defer changeStream.Close(context.Background())

	for changeStream.Next(context.Background()) {
		var changeDocument bson.M
		if err = changeStream.Decode(&changeDocument); err != nil {
			panic(err)
		}

		parseChangeDocument(changeDocument)
	}
}

func parseChangeDocument(changeDocument bson.M) {
	changeBytes, err := json.Marshal(changeDocument)
	if err != nil {
		log.Println(err)
	}

	var parsedDocument map[string]interface{}
	if err = json.Unmarshal(changeBytes, &parsedDocument); err != nil {
		log.Println(err)
	}

	operationType := parsedDocument["operationType"]
	data := parsedDocument["fullDocument"]
	collection := parsedDocument["ns"].(map[string]interface{})["coll"].(string)
	documentKey := parsedDocument["documentKey"].(map[string]interface{})["_id"].(string)

	if operationType == "delete" {
		fmt.Println("Delete ", documentKey, " from ", collection, " collection")
	} else {
		if value, ok := data.(map[string]interface{}); ok {
			// flow.Run(collection, operationType, value)

			// escape
			log.Println(value)
		}
	}
}
