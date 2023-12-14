package database

import (
	"fmt"
	"log"
	"time"
	"os"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"github.com/joho/godotenv"
)

func DBConnect() *mongo.Client {
	err := godotenv.Load(".env")

	if err != nil{
		log.Fatal("Err Loading .env file")
	}

	Mongo_DB := os.Getenv("MONGO_URI")

	client, error := mongo.NewClient(options.Client().ApplyURI(Mongo_DB))

        if error != nil{
			log.Fatal(error)
		}
}