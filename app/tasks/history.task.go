package tasks

import (
	"time"

	"example.com/refcode/v1/app/models"
	"github.com/RichardKnop/machinery/v1/log"
)

func SaveHistory(address, domain, created string, amount float64) error {
	timeFormat := "2006-01-02T15:04:05.999Z"
	_created, err := time.Parse(timeFormat, created)
	if err != nil {
		return err
	}
	log.DEBUG.Println("WITHDRAW EVENTS")
	log.DEBUG.Printf("TO ADDRESS %s\n", address)
	log.DEBUG.Printf("AMOUNT %f\n", amount)
	log.DEBUG.Printf("DOMAIN %s\n", domain)
	history := models.WithdrawHistory{
		Address: address,
		Amount:  amount,
		Created: _created,
		Domain:  domain,
	}
	log.DEBUG.Println("HISTORY SAVE OK")
	return history.Save()
}
