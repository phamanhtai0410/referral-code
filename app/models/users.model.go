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
	TableDetails string = "user-detail"
)

type User struct {
	Created           time.Time `json:"created"`
	RefCode           int64     `json:"refcode"`
	Address           string    `json:"address"`
	Counter           int64     `json:"counter"`
	Level             string    `json:"level"`
	Rate              float64   `json:"rate"`
	WithdrawAvailable float64   `json:"withdraw_available" bson:"withdraw_available"`
}

func (ref *User) Save() error {
	collection := database.GetCollection(TableDetails)
	ctx, _ := database.NewContext()
	_, err := collection.InsertOne(ctx, *ref)
	return err
}

func (ref *User) FindByAddress(address string) (int64, error) {
	var result User
	collections := database.GetCollection(TableDetails)
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

func (ref *User) FindDocsByAddress(address string) (*User, error) {
	var result = new(User)
	collections := database.GetCollection(TableDetails)
	ctx, _ := database.NewContext()
	filter := bson.D{{"address", address}}
	err := collections.FindOne(ctx, filter).Decode(result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, err
		}
		log.Fatal("[MODEL] ", err)
	}
	return result, nil
}

func (ref *User) IsExits(code string) (int64, bool) {
	var result User
	collection := database.GetCollection(TableDetails)
	filter := bson.D{{"refcode", code}}
	if err := collection.FindOne(context.TODO(), filter).Decode(&result); err != nil {
		if err == mongo.ErrNoDocuments {
			return -1, false
		}
	}
	return result.Counter, true
}

func (ref *User) UpdateCounter(counter int64) error {
	collection := database.GetCollection(TableDetails)
	filter := bson.M{"refcode": ref.RefCode}
	update := bson.M{
		"$set": bson.M{
			"counter": counter,
		},
	}
	_, err := collection.UpdateOne(context.TODO(), filter, update)
	return err
}

func (ref *User) UpdateRecord(
	refcode, counter int64,
	withdrawAvailable, rate float64,
	level string,
) error {
	collection := database.GetCollection(TableDetails)
	filter := bson.M{"refcode": refcode}
	update := bson.M{
		"$set": bson.M{
			"counter":            counter,
			"level":              level,
			"rate":               rate,
			"withdraw_available": withdrawAvailable,
		},
	}
	_, err := collection.UpdateOne(context.TODO(), filter, update)
	return err
}

func (ref *User) GetAllRecords() ([]User, error) {
	collection := database.GetCollection(TableDetails)
	var codes []User

	cursor, err := collection.Find(context.TODO(), bson.D{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	for cursor.Next(context.TODO()) {
		var code User
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
