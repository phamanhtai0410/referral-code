package services

import (
	"encoding/json"
	"log"
	"time"

	"example.com/refcode/v1/app/models"
	"example.com/refcode/v1/platform/cache"

	"example.com/refcode/v1/app/schemas"
	"example.com/refcode/v1/pkg/constants"
	"example.com/refcode/v1/platform/queue"
)

func GenRefCode(address string) int64 {
	var refcode models.RefCode
	code, _ := refcode.FindByAddress(address)
	if code == -1 {
		val, err := cache.Incr(constants.CacheCounter)
		if err != nil {
			log.Fatal("[SERVICE] " + err.Error())
		}
		code = val
		// SaveRefCode(&schemas.RefCodeRequest{
		// 	Address: address,
		// 	RefCode: val,
		// })
		model := models.RefCode{
			RefCode: val,
			Address: address,
			Created: time.Now().UTC(),
			Counter: 0,
		}
		err = model.Save()
		if err != nil {
			log.Fatal(constants.LogWorkerSaveCodeDetails, err)
		}
	}
	return code
}

func SaveRefCodeInfo(req *schemas.RefCodeUsedRequest) {
	_, err := queue.NewQueue(constants.WorkerQueue)
	if err != nil {
		log.Fatal("failed to create queue " + err.Error())
	}
	data, err := json.Marshal(req)
	if err != nil {
		log.Fatal("failed to encode " + err.Error())
	}
	if err = queue.Publish(
		constants.WorkerQueue,
		constants.MsgSaveRefCodeUsed,
		data,
	); err != nil {
		log.Fatal("Failed to publish msg " + err.Error())
	}
	log.Println("[SERVICES] Publishing to queue ...")
}

func SaveRefCode(req *schemas.RefCodeRequest) {
	_, err := queue.NewQueue(constants.WorkerQueue)
	if err != nil {
		log.Fatal("failed to create queue " + err.Error())
	}
	data, err := json.Marshal(req)
	if err != nil {
		log.Fatal("failed to encode " + err.Error())
	}
	if err = queue.Publish(
		constants.WorkerQueue,
		constants.MsgSaveRefCodeDetail,
		data,
	); err != nil {
		log.Fatal("Failed to publish msg " + err.Error())
	}
	log.Println("[SERVICES] Publishing to queue ...")
}
