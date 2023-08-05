package tasks

import (
	"log"
	"time"

	"example.com/refcode/v1/app/models"
	"example.com/refcode/v1/pkg/constants"
	"example.com/refcode/v1/pkg/utils"
	"example.com/refcode/v1/pkg/workers"
)

func RefCodeCounterUsed() {
	ticker := time.NewTicker(constants.JobSchedule * time.Minute)
	defer ticker.Stop()
	CountNumberOfRefCodeUsed(<-ticker.C)
	for _time := range ticker.C {
		go CountNumberOfRefCodeUsed(_time)
	}
}

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
			err := workers.Delay(
				"Worker.SaveUserInfo",
				SaveUserInfo,
				time.Now().UTC().Format("2006-01-02T15:04:05.999Z"),
				time.Now().UTC().Format("2006-01-02T15:04:05.999Z"),
				key,
				level,
				val,
				rate,
				float64(0),
				price[key]*rate,
			)
			if err != nil {
				log.Fatal(err)
			}
		} else {
			level, rate = utils.ReferralRule(int(val + user.Counter))
			err := workers.Delay(
				"Worker.UpdateUserRecord",
				UpdateUserRecord,
				key,
				val+user.Counter,
				user.TotalEarn,
				rate,
				(price[key]*rate)+user.Pending,
				level,
			)
			if err != nil {
				log.Fatal(err)
			}
		}
		log.Printf("referral code %s count: %d\n", key, val)
	}
}
