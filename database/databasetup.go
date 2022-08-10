package database

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

func DBSet() *mongo.Client {

	//serverAPIOptions := options.ServerAPI(options.ServerAPIVersion1)
	//clientOptions := options.Client().ApplyURI("mongodb+srv://mongo5m0k13:<password>@cluster0.u1esg.mongodb.net/?retryWrites=true&w=majority").SetServerAPIOptions(serverAPIOptions)
	//ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	//defer cancel()
	//client, err := mongo.Connect(ctx, clientOptions)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//mongodb: //mongodb:27017
	//client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://development:testpassword@localhost:27017"))

	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://mongo5m0k13:eknmgptctkpr@cluster0.npyvkdk.mongodb.net/?retryWrites=true&w=majority"))

	//
	//client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://mongo5m0k13:eknmgptctkpr@cluster0.u1esg.mongodb.net/?retryWrites=true&w=majority").SetServerAPIOptions(serverAPIOptions))

	//client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://mongo5m0k13:eknmgptctkpr@cluster0.u1esg.mongodb.net/?retryWrites=true&w=majority"))
	if err != nil {
		log.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		log.Println("!!!failed to connect to mongodb!!!!!!!")
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Println("failed to connect to mongodb!!!!!!!")
		return nil
	}
	println("Successfully Connected to the mongodb")
	return client
}

var Client *mongo.Client = DBSet()

// for main.go
func UserData(client *mongo.Client, collectionName string) *mongo.Collection {
	var collection *mongo.Collection = client.Database("Ecommerce").Collection(collectionName)
	return collection
}

func ProductData(client *mongo.Client, collectionName string) *mongo.Collection {
	var ProductCollection *mongo.Collection = client.Database("Ecommerce").Collection(collectionName)
	return ProductCollection
}
