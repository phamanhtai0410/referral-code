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
	LastUpdated       time.Time `json:"last_updated" bson:"last_updated"`
	Id                int64     `json:"id"`
	ReferralCode      string    `json:"referral_code" bson:"referral_code"`
	Address           string    `json:"address"`
	Counter           int64     `json:"counter"`
	Level             string    `json:"level"`
	Rate              float64   `json:"rate"`
	WithdrawAvailable float64   `json:"withdraw_available" bson:"withdraw_available"`
}

func (ref *User) Save() error {
	ref.LastUpdated = time.Now().UTC()
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
	return result.Id, nil
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

func (ref *User) OwnerOf(domain string) (string, int64) {
	var result User
	collection := database.GetCollection(TableDetails)
	filter := bson.M{
		"referral_code": domain,
	}
	if err := collection.FindOne(context.TODO(), filter).Decode(&result); err != nil {
		if err == mongo.ErrNoDocuments {
			return "", -1
		}
	}
	return result.Address, result.Id
}

func (ref *User) IsExits(address, domain string) bool {
	var result User
	collection := database.GetCollection(TableDetails)
	filter := bson.M{
		"address":       address,
		"referral_code": domain,
	}
	if err := collection.FindOne(context.TODO(), filter).Decode(&result); err != nil {
		if err == mongo.ErrNoDocuments {
			return false
		}
	}
	return true
}

func (ref *User) WalletExists(address string) bool {
	var result User
	collection := database.GetCollection(TableDetails)
	filter := bson.M{
		"address": address,
	}
	if err := collection.FindOne(context.TODO(), filter).Decode(&result); err != nil {
		if err == mongo.ErrNoDocuments {
			return false
		}
	}
	return true
}

func (ref *User) UpdateCounter(counter int64) error {
	collection := database.GetCollection(TableDetails)
	filter := bson.M{"id": ref.Id}
	update := bson.M{
		"$set": bson.M{
			"counter": counter,
		},
	}
	_, err := collection.UpdateOne(context.TODO(), filter, update)
	return err
}

func (ref *User) UpdateRecord(
	id, counter int64,
	withdrawAvailable, rate float64,
	level string,
) error {
	collection := database.GetCollection(TableDetails)
	filter := bson.M{"id": id}
	update := bson.M{
		"$set": bson.M{
			"counter":            counter,
			"level":              level,
			"rate":               rate,
			"withdraw_available": withdrawAvailable,
			"last_updated":       time.Now().UTC(),
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
