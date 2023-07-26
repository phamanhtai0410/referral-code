package workers

import (
	"example.com/refcode/v1/pkg/constants"
	"fmt"
	"log"
	"time"

	"example.com/refcode/v1/app/models"
)

func CountNumberOfRefCodeUsed(_time time.Time) {
	currentTime := _time
	oneMinuteAgo := _time.Add(-1 * constants.JobSchedule * time.Minute)
	tableRefCode := new(models.RefCode)
	tableCodeUsed := new(models.RefCodeUsed)
	codes, err := tableRefCode.GetAllRecords()
	if err != nil {
		log.Fatal("[WORKER #1] count number of ref codes ", err)
	}
	for _, code := range codes {
		tableCodeUsed.RefCode = code.RefCode
		tableRefCode.RefCode = code.RefCode
		count, _ := tableCodeUsed.CountDocumentsByTime(
			oneMinuteAgo.UTC(),
			currentTime.UTC(),
		)
		if count != 0 {
			err = tableRefCode.UpdateCounter(code.Counter + count)
			if err != nil {
				return
			}
			logMsg := fmt.Sprintf("[WORKER #2] count number of ref codes %d: %d", code.RefCode, code.Counter+count)
			log.Println(logMsg)
		}
	}
}
