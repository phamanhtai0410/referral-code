package models

import (
	"context"
	"log"
	"time"

	"example.com/refcode/v1/platform/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	tableDetails string = "code-detail"
)

type RefCode struct {
	Created time.Time `json:"created"`
	RefCode int64     `json:"refcode"`
	Address string    `json:"address"`
	Counter int64     `json:"counter"`
}

func (ref *RefCode) Save() error {
	collection := database.GetCollection(tableDetails)
	ctx, _ := database.NewContext()
	_, err := collection.InsertOne(ctx, *ref)
	return err
}

func (ref *RefCode) FindByAddress(address string) (int64, error) {
	var result RefCode
	collections := database.GetCollection(tableDetails)
	ctx, _ := database.NewContext()
	filter := bson.D{{"address", address}}
	err := collections.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return -1, err
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

func (ref *RefCode) GetAllRecords() ([]RefCode, error) {
	collection := database.GetCollection(tableDetails)
	var codes []RefCode

	cursor, err := collection.Find(context.TODO(), bson.D{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	for cursor.Next(context.TODO()) {
		var code RefCode
		if err := cursor.Decode(&code); err != nil {
			return nil, err
		}
		codes = append(codes, code)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return codes, nil
}
