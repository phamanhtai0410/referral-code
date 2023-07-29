package models

import (
	"context"
	"errors"
	"example.com/refcode/v1/platform/database"
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

const (
	tableHistory string = "withdraw-history"
)

type WithdrawHistory struct {
	Address string    `json:"address"`
	Domain  string    `json:"domain"`
	Amount  float64   `json:"amount"`
	Created time.Time `json:"created"`
}

func (h *WithdrawHistory) Save() error {
	collection := database.GetCollection(tableHistory)
	_, err := collection.InsertOne(context.TODO(), *h)
	return err
}

func (h *WithdrawHistory) GetDocumentByDomain(domain string) ([]WithdrawHistory, error) {
	collection := database.GetCollection(tableHistory)
	var histories []WithdrawHistory

	cursor, err := collection.Find(context.TODO(), bson.M{"domain": domain})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	for cursor.Next(context.TODO()) {
		var history WithdrawHistory
		if err := cursor.Decode(&history); err != nil {
			return nil, err
		}
		histories = append(histories, history)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}
	if len(histories) == 0 {
		return nil, errors.New("no data")
	}
	return histories, nil
}
