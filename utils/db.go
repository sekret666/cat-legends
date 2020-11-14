package utils

import (
	"context"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
)

var dataBase DataBase

func GetDB() *DataBase {
	return &dataBase
}

type DataBase struct {
	Ctx     context.Context
	client  *mongo.Client
	Players *mongo.Collection
}

func InitDB() {
	var err error
	ctx := context.Background()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("MONGODB")))
	if err != nil {
		log.Fatal(err)
	}

	playersCollection := client.Database("cat_legends").Collection("players")

	dataBase = DataBase{
		Ctx:     ctx,
		client:  client,
		Players: playersCollection,
	}

	log.Info("Database init successful")
}

func CloseDB() {
	if err := dataBase.client.Disconnect(dataBase.Ctx); err != nil {
		log.Fatal(err)
	}
	log.Info("Database closed successfully")
}
