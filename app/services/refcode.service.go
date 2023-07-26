package services

import (
	"encoding/json"
	"errors"
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
			log.Fatal("[SERVICE #2] " + err.Error())
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

func SaveRefCodeInfo(req *schemas.RefCodeUsedRequest) error {
	currentCode, err := cache.Get(constants.CacheCounter)
	if err != nil {
		return err
	}
	if req.RefCode <= constants.CacheCounterBegin || req.RefCode > currentCode {
		return errors.New("referral code does not exist yet")
	}
	_, err = queue.NewQueue(constants.WorkerQueue)
	if err != nil {
		log.Fatal("[SERVICE #2] failed to create queue " + err.Error())
	}
	data, err := json.Marshal(req)
	if err != nil {
		log.Fatal("[SERVICE #3] failed to encode " + err.Error())
	}
	if err = queue.Publish(
		constants.WorkerQueue,
		constants.MsgSaveRefCodeUsed,
		data,
	); err != nil {
		log.Fatal("[SERVICE #4] Failed to publish msg " + err.Error())
	}
	log.Println("[SERVICES #5] Publishing to queue ...")
	return nil
}

func SaveRefCode(req *schemas.RefCodeRequest) {
	_, err := queue.NewQueue(constants.WorkerQueue)
	if err != nil {
		log.Fatal("[SERVICE #6] failed to create queue " + err.Error())
	}
	data, err := json.Marshal(req)
	if err != nil {
		log.Fatal("[SERVICE #7] failed to encode " + err.Error())
	}
	if err = queue.Publish(
		constants.WorkerQueue,
		constants.MsgSaveRefCodeDetail,
		data,
	); err != nil {
		log.Fatal("[SERVICE #8] Failed to publish msg " + err.Error())
	}
	log.Println("[SERVICES #9] Publishing to queue ...")
}

func RefCodeTracking(address string) (*schemas.TrackingResponse, error) {
	model := models.RefCode{
		Address: address,
	}
	code, err := model.FindDocsByAddress(address)
	if err != nil {
		return nil, err
	}
	resp := new(schemas.TrackingResponse)
	if code.Counter > 0 && code.Counter <= 30 {
		resp.Rate = 0.1
		resp.Level = "Standard"
	} else if code.Counter >= 31 && code.Counter <= 100 {
		resp.Rate = 0.2
		resp.Level = "Level I"
	} else if code.Counter >= 101 && code.Counter <= 1000 {
		resp.Rate = 0.3
		resp.Level = "Level II"
	} else if code.Counter >= 1001 && code.Counter <= 5000 {
		resp.Rate = 0.5
		resp.Level = "Level III"
	} else {
		resp.Rate = 0.7
		resp.Level = "Level IV"
	}
	resp.Count = code.Counter
	return resp, nil
}
