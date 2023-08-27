package mongo_db

import (
	"context"
	"cycling-tracker-server/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

var mongoDBInstance *mongo.Client = nil

func ConnectToMongoDB() {
	if mongoDBInstance != nil {
		return
	}

	newMongoInstance, err := mongo.NewClient(options.Client().ApplyURI(config.GetENVByKey("MONGODB_URL")))
	if err != nil {
		log.Printf("There was an error creating the mongoDB mongoDBInstance: %v", err)
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = newMongoInstance.Connect(ctx)

	if err != nil {
		log.Printf("There was an error connecting to the mongoDB with mongoDBInstance: %v", err)
	}

	err = newMongoInstance.Ping(ctx, nil)
	if err != nil {
		log.Printf("There was an error pinging the mongodb database: %v", err)
	}

	log.Printf("Successfuly Connected to MongoDB")

	mongoDBInstance = newMongoInstance
}

func GetMongoDB() *mongo.Client {
	if mongoDBInstance == nil {
		ConnectToMongoDB()
	}

	return mongoDBInstance
}

func GetCollection(collection string) *mongo.Collection {
	return GetMongoDB().Database("cycling-tracker-db").Collection(collection)
}
