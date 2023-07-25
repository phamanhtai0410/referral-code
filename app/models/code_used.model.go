package models

import (
	"example.com/refcode/v1/platform/database"
)

const (
	tableCodeUsed string = "code-used"
)

type RefCodeUsed struct {
	RefCode string  `json:"refcode"`
	Domain  string  `json:"domain"`
	Price   float64 `json:"price"`
	Address string  `json:"address"`
	Created string  `bson:"created"`
}

func (r *RefCodeUsed) Save() error {
	collection := database.GetCollection(tableCodeUsed)
	ctx, _ := database.NewContext()
	_, err := collection.InsertOne(ctx, *r)
	return err
}
