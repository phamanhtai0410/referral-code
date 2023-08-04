package models

import (
	"context"
	"time"

	"example.com/refcode/v1/platform/database"
	"go.mongodb.org/mongo-driver/bson"
)

const (
	tableCodeUsed string = "referral-code-used"
)

type CodeUsed struct {
	ReferralCode string    `json:"referral_code" bson:"referral_code"`
	Domain       string    `json:"domain"`
	Price        float64   `json:"price"`
	Address      string    `json:"address"`
	Created      time.Time `bson:"created"`
}

func (r *CodeUsed) Save() error {
	collection := database.GetCollection(tableCodeUsed)
	ctx, _ := database.NewContext()
	_, err := collection.InsertOne(ctx, *r)
	return err
}

func (r *CodeUsed) GetDocumentsByTime(startTime, endTime time.Time) ([]CodeUsed, error) {
	collection := database.GetCollection(tableCodeUsed)
	filter := bson.M{
		"created": bson.M{
			"$gte": startTime,
			"$lte": endTime,
		},
	}

	var records []CodeUsed
	cursor, err := collection.Find(context.TODO(), filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	for cursor.Next(context.TODO()) {
		var record CodeUsed
		if err = cursor.Decode(&record); err != nil {
			return nil, err
		}
		records = append(records, record)
	}

	if err = cursor.Err(); err != nil {
		return nil, err
	}

	return records, nil
}
