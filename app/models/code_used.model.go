package models

import (
	"context"
	"time"

	"example.com/refcode/v1/platform/database"
	"go.mongodb.org/mongo-driver/bson"
)

const (
	tableCodeUsed string = "code-used"
)

type RefCodeUsed struct {
	RefCode int64  `json:"refcode"`
	Domain  string  `json:"domain"`
	Price   float64 `json:"price"`
	Address string  `json:"address"`
	Created time.Time  `bson:"created"`
}

func (r *RefCodeUsed) Save() error {
	collection := database.GetCollection(tableCodeUsed)
	ctx, _ := database.NewContext()
	_, err := collection.InsertOne(ctx, *r)
	return err
}

func (r *RefCodeUsed) CountDocumentsByTime(startTime, endTime time.Time) (int64, error) {
	collection := database.GetCollection(tableCodeUsed)
	filter := bson.M{
		"refcode": r.RefCode,
		"created": bson.M{
			"$gte": startTime,
			"$lte": endTime,
		},
	}

	count, err := collection.CountDocuments(context.TODO(), filter)
	if err != nil {
		return 0, err
	}

	return count, nil
}
