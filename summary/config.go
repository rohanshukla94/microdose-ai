package summary

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
	"time"
)

func MongoURI() string {

	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading MongoURI string")
	}
	return os.Getenv("MONGO_URI")
}

func pt3ApiKey() string {

	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading openai api key")
	}
	return os.Getenv("OPENAI_KEY")
}

func ConnectDB() *mongo.Client {
	client, err := mongo.NewClient(options.Client().ApplyURI(MongoURI()))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	//ping the database
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB")
	return client
}

var DB *mongo.Client = ConnectDB()

func GetCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	collection := client.Database("summary_ai").Collection(collectionName)
	return collection
}
