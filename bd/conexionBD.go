package bd

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)
var MongoCN = ConnectDB()
var clientOptions = options.Client().ApplyURI("mongodb+srv://jsosa:f5Z0zOIyThQapFmc@cluster0.7g5qi.mongodb.net/twittor?retryWrites=true&w=majority")
/*ConnectDB Create the Database connection*/
func ConnectDB() *mongo.Client{
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}
	log.Println("Success Connection ")
	return client
}

/*CheckConnection make a ping to the the Database connection*/
func CheckConnection() int{
	err := MongoCN.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err.Error())
		return 0
	}
	return 1
}