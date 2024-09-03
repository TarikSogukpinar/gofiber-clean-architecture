package database

import (
	"context"
	"fmt"
	"gofiber-clean-architecture/configuration"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var UserCollection *mongo.Collection
var BookCollection *mongo.Collection

type MongoInstance struct {
	Client *mongo.Client
	Db     *mongo.Database
}

var Mg MongoInstance

func Connect() error {

	configuration.LoadConfig()

	mongoURI := configuration.Get("MONGODB_URI")
	dbName := configuration.Get("DB_NAME")

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(mongoURI).SetMaxPoolSize(50).SetConnectTimeout(10*time.Second))

	if err != nil {
		return err
	}

	// Ping the database to verify connection
	if err := client.Ping(context.TODO(), nil); err != nil {
		return err
	}

	db := client.Database(dbName)

	Mg = MongoInstance{
		Client: client,
		Db:     db,
	}

	UserCollection = db.Collection("users")
	// BookCollection = db.Collection("books")

	fmt.Println("Connected to MongoDB!")
	return nil
}
