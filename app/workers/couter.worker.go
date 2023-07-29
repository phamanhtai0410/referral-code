package workers

import (
	"example.com/refcode/v1/pkg/utils"
	"log"
	"time"

	"example.com/refcode/v1/app/models"

	"example.com/refcode/v1/pkg/constants"
)

func CountNumberOfRefCodeUsed(_time time.Time) {
	var codeUsedModel models.CodeUsed
	var userModel models.User
	var count = make(map[string]int64)
	var price = make(map[string]float64)

	currentTime := _time
	oneMinuteAgo := _time.Add(-1 * constants.JobSchedule * time.Minute)

	codeUsedArr, err := codeUsedModel.GetDocumentsByTime(
		oneMinuteAgo.UTC(),
		currentTime.UTC(),
	)
	if err != nil {
		log.Fatal("[WORKER] " + err.Error())
	}
	for _, code := range codeUsedArr {
		count[code.ReferralCode]++
		price[code.ReferralCode] += code.Price

	}
	for key, val := range count {
		user, errFind := userModel.FindDocsByReferralCode(key)
		level, rate := utils.ReferralRule(int(val))
		if errFind != nil { // No documents
			newUser := models.User{
				ReferralCode: key,
				Counter:      val,
				Level:        level,
				Rate:         rate,
				Created:      time.Now().UTC(),
				LastUpdated:  time.Now().UTC(),
				TotalEarn:    price[key] * rate,
			}
			err = newUser.Save()
			if err != nil {
				log.Fatal("[WORKER] " + err.Error())
			}
		} else {
			if level == "" {
				level = user.Level
				rate = user.Rate
			}
			err = user.UpdateRecord(key, val+user.Counter, (price[key]*rate)+user.TotalEarn, rate, level)
			if err != nil {
				log.Fatal("[WORKER] " + err.Error())
			}
		}
		log.Printf("referral code %s count: %d\n", key, val)
	}
}
