package links

import (
	"context"
	"log"
	"time"

	"github.com/Salomon-Novachrono/graphQL-test/db/modles/users"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Link struct {
	ID      string
	Title   string
	Address string
	User    *users.User
}

func SaveLink(link Link, dbCon mongo.Database) {

	data := bson.D{{"_id", link.ID}, {"title", link.Title}, {"Address", link.Address}, {"User", link.User.Username}}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	collcetion := dbCon.Collection("Links")
	res, err := collcetion.InsertOne(ctx, data)
	if err != nil {
		panic(err)
	}
	log.Println("created new user with id: {}", res.InsertedID)
}
