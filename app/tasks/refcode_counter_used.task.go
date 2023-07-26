package tasks

import (
	"example.com/refcode/v1/pkg/constants"
	"time"

	"example.com/refcode/v1/app/workers"
)

func RefCodeCounterUsed() {
	ticker := time.NewTicker(constants.JobSchedule * time.Minute)
	defer ticker.Stop()
	workers.CountNumberOfRefCodeUsed(<-ticker.C)
	for _time := range ticker.C {
		go workers.CountNumberOfRefCodeUsed(_time)
	}
}
