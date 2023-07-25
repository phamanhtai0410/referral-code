package models

import (
	"context"
	"log"

	"example.com/refcode/v1/platform/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	tableDetails string = "code-detail"
)

type RefCode struct {
	Created string `json:"created"`
	RefCode string `json:"refcode"`
	Address string `json:"address"`
	Counter int64  `json:"counter"`
}

func (ref *RefCode) Save() error {
	collection := database.GetCollection(tableDetails)
	ctx, _ := database.NewContext()
	_, err := collection.InsertOne(ctx, *ref)
	return err
}

func (ref *RefCode) FindByAddress(address string) (string, error) {
	var result RefCode
	collections := database.GetCollection(tableDetails)
	ctx, _ := database.NewContext()
	filter := bson.D{{"address", address}}
	err := collections.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return "", err
		}
		log.Fatal("[MODEL] ", err)
	}
	return result.RefCode, nil
}

func (ref *RefCode) IsExits(code string) (int64, bool) {
	var result RefCode
	collection := database.GetCollection(tableDetails)
	filter := bson.D{{"refcode", code}}
	if err := collection.FindOne(context.TODO(), filter).Decode(&result); err != nil {
		if err == mongo.ErrNoDocuments {
			return -1, false
		}
	}
	return result.Counter, true
}

func (ref *RefCode) UpdateCounter(counter int64) error {
	collection := database.GetCollection(tableDetails)
	filter := bson.M{"refcode": ref.RefCode}
	update := bson.M{
		"$set": bson.M{
			"counter": counter,
		},
	}
	_, err := collection.UpdateOne(context.TODO(), filter, update)
	return err
}
