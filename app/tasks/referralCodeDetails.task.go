package tasks

import (
	"time"

	"example.com/refcode/v1/app/models"
	"github.com/RichardKnop/machinery/v1/log"
)

func UpdateUserRecord(refCodeCondition string, counter int64, totalEarn float64, rate float64, level string) error {
	var user models.User
	err := user.UpdateRecord(refCodeCondition, counter, totalEarn, rate, level)
	if err != nil {
		return err
	}
	log.DEBUG.Println("Update OK")
	return nil
}

func SaveUserInfo(created, uploaded, referralCode, level string, counter int64, rate, totalEarn float64) error {
	timeFormat := "2006-01-02T15:04:05.999Z"
	_created, err := time.Parse(timeFormat, created)
	if err != nil {
		return err
	}
	_uploaded, err := time.Parse(timeFormat, uploaded)
	if err != nil {
		return err
	}
	log.INFO.Println()
	newUser := models.User{
		ReferralCode: referralCode,
		Counter:      counter,
		Level:        level,
		Rate:         rate,
		Created:      _created,
		LastUpdated:  _uploaded,
		TotalEarn:    totalEarn,
	}
	log.DEBUG.Println("SAVE OK")
	return newUser.Save()
}
