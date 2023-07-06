package config

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

//https://cloud.mongodb.com/ access with meli account
//mongodb+srv://nico:<password>@cluster0.av039p4.mongodb.net/?retryWrites=true&w=majority

const (
	cnnString = "mongodb+srv://nico:Martoneto_365@cluster0.av039p4.mongodb.net/?retryWrites=true&w=majority"
)

func CreateMongoDBClient() (client *mongo.Client)  {
	var ctx = context.TODO()
	clientOptions := options.Client().ApplyURI(cnnString)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	return client
}
