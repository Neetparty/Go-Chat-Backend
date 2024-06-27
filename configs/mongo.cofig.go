package configs

import (
	"fmt"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Mongo *mongo.Client

// CreateCollection creates collections in the MongoDB database.
func CreateCollection(names []string) error {
	db := Mongo.Database("go-chat")

	for _, name := range names {
		coll := db.Collection(name)
		if coll == nil {
			if err := db.CreateCollection(nil, name); err != nil {
				return err
			} else {
				fmt.Println("Collection " + name + " created")
			}
		} else {
			fmt.Println("Collection " + name + " already exists")
		}
	}

	return nil
}

// InitMongo initializes the MongoDB client.
func InitMongo() error {
	mongoURI := os.Getenv("MONGO_URL")
	client, err := mongo.Connect(nil, options.Client().ApplyURI(mongoURI))
	if err != nil {
		return err
	}
	Mongo = client

	collections := []string{
		"user",
		"room",
		"message",
	}

	err = CreateCollection(collections)
	if err != nil {
		return err
	}

	return nil
}
