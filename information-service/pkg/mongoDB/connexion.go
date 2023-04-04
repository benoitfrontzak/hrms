package mongoDB

import (
	"context"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client
var DSN = os.Getenv("DSN")

const DB = "logs"

// ConnectToDB connect to mongo db
func ConnectToDB() (*mongo.Client, error) {
	// create connection options
	clientOptions := options.Client().ApplyURI(DSN)
	// clientOptions.SetAuth(options.Credential{
	// 	Username: "admin",
	// 	Password: "password",
	// })

	// connect
	c, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Println("Error connecting:", err)
		return nil, err
	}

	log.Println("Connected to mongo!")

	return c, nil
}
