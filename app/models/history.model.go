package models

import (
	"context"
	"example.com/refcode/v1/platform/database"
	"time"
)

const (
	tableHistory string = "withdraw-history"
)

type WithdrawHistory struct {
	Address string    `json:"address"`
	Amount  float64   `json:"amount"`
	Created time.Time `json:"created"`
}

func (h *WithdrawHistory) Save() error {
	collection := database.GetCollection(tableHistory)
	_, err := collection.InsertOne(context.TODO(), *h)
	return err
}
