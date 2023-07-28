package workers

import (
	"fmt"
	"log"
	"time"

	"example.com/refcode/v1/pkg/constants"
	"example.com/refcode/v1/pkg/utils"

	"example.com/refcode/v1/app/models"
)

func CountNumberOfRefCodeUsed(_time time.Time) {
	currentTime := _time
	oneMinuteAgo := _time.Add(-1 * constants.JobSchedule * time.Minute)
	userModel := new(models.User)
	codeUsedModel := new(models.CodeUsed)
	users, err := userModel.GetAllRecords()
	if err != nil {
		log.Fatal("[WORKER #1] count number of ref codes ", err)
	}
	for _, user := range users {
		codeUsedModel.Id = user.Id
		records, err := codeUsedModel.GetDocumentsByTime(
			oneMinuteAgo.UTC(),
			currentTime.UTC(),
		)
		if err != nil {
			log.Fatal("[WORKER #2]", err)
		}
		count := int64(len(records))
		level, rate := utils.ReferralRule(len(records))
		if level == "" {
			level = user.Level
			rate = user.Rate
		}
		//log.Println(count)
		var totalPrice float64 = 0
		for _, record := range records {
			totalPrice += record.Price
		}
		err = userModel.UpdateRecord(
			user.Id,
			user.Counter+count,
			rate*totalPrice+user.WithdrawAvailable,
			rate,
			level,
		)
		if err != nil {
			log.Fatal("[WORKER #3]", err)
		}
		logMsg := fmt.Sprintf("[WORKER #4] count number of ref codes %d: %d", user.Id, user.Counter+count)
		log.Println(logMsg)
	}
}
