package workers

import (
	"fmt"
	"log"
	"time"

	"example.com/refcode/v1/app/models"
)

func CountNumberOfRefCodeUsed(_time time.Time) {
	currentTime := _time
	oneMinuteAgo := _time.Add(-1 * time.Minute)
	tableRefCode := new(models.RefCode)
	tableCodeUsed := new(models.RefCodeUsed)
	codes, err := tableRefCode.GetAllRecords()
	if err != nil {
		log.Fatal("[WORKER] count number of ref codes ", err)
	}
	for _, code := range codes {
		tableCodeUsed.RefCode = code.RefCode
		tableRefCode.RefCode = code.RefCode
		count, _ := tableCodeUsed.CountDocumentsByTime(
			oneMinuteAgo.UTC(),
			currentTime.UTC(),
		)
		if count != 0 {
			tableRefCode.UpdateCounter(code.Counter + count)
			logMsg := fmt.Sprintf("[WORKER JOBS] count number of ref codes %d: %d", code.RefCode, code.Counter+count)
			log.Println(logMsg)
		}
	}
}
