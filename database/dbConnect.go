package database

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

func DBConnect() *mongo.Client {
	err := godotenv.Load(".env")

	if err != nil{
		log.Fatalln("Err Loading .env file")
	}

	Mongo_DB := os.Getenv("MONGO_URI")

	clientOptions := options.Client().ApplyURI(Mongo_DB)

	client, err := mongo.NewClient(clientOptions)

        if err != nil{
			log.Fatalln(err)
		}

		ctx, cancel :=  context.WithTimeout(context.Background(), 10*time.Second)

		defer cancel()

	 client.Connect(ctx)

		if err != nil{
			log.Fatalln(err)
		}

		fmt.Println("Database Connected and We are up people")

		return client

}


var Client *mongo.Client = DBConnect()

func NewCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	collection := client.Database("clluster0").Collection(collectionName)
	return collection
}
