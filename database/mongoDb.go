package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/Salomon-Novachrono/graphQL-test/graph/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type DB struct {
	client *mongo.Client
}

func Connect(dbUrl string) *DB {
	client, err := mongo.NewClient(options.Client().ApplyURI(dbUrl))
	if err != nil {
		log.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel = context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}

	return &DB{
		client: client,
	}

}

func (db *DB) InsertHumanById(human model.NewHuman) *model.Human {
	humanColl := db.client.Database("graphQl-db").Collection("human")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	inserg, err := humanColl.InsertOne(ctx, bson.D{{Key: "name", Value: human.Name}})
	if err != nil {
		log.Fatal(err)
	}
	insertedID := inserg.InsertedID.(primitive.ObjectID).Hex()
	returnHuman := model.Human{ID: insertedID, Name: human.Name}

	return &returnHuman
}

func (db *DB) FindHumanById(id string) *model.Human {
	ObjectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Fatal(err)
	}
	humanColl := db.client.Database("graphQl-db").Collection("human")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	res := humanColl.FindOne(ctx, bson.M{"_id": ObjectID})

	human := model.Human{ID: id}

	res.Decode(&human)

	return &human
}

func (db *DB) All() []*model.Human {
	humanColl := db.client.Database("graphQl-db").Collection("human")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	cur, err := humanColl.Find(ctx, bson.D{})
	if err != nil {
		log.Fatal(err)
	}
	var humans []*model.Human
	for cur.Next(ctx) {

		sus, err := cur.Current.Elements()
		fmt.Println(sus)
		if err != nil {
			log.Fatal(err)
		}
		human := model.Human{ID: (""), Name: ("")}

		humans = append(humans, &human)
	}
	return humans
}
