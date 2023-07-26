package services

import (
	"example.com/refcode/v1/pkg/constants"
	"example.com/refcode/v1/platform/queue"
	"log"
)

func RefCodeCounter(code string) {
	if err := queue.Publish(
		constants.WorkerQueue,
		constants.MsgRefCodeCounter,
		[]byte(code),
	); err != nil {
		log.Fatal("[COUNTER #1] Failed to publish msg " + err.Error())
	}
	log.Println("[SERVICES #1] Publishing to queue ...")
}
