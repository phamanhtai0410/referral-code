package tasks

import (
	"time"

	"example.com/refcode/v1/app/workers"
)

func RefCodeCounterUsed() {
	ticker := time.NewTicker(5 * time.Minute)
	defer ticker.Stop()
	workers.CountNumberOfRefCodeUsed(<-ticker.C)
	for _time := range ticker.C {
		go workers.CountNumberOfRefCodeUsed(_time)
	}
}
