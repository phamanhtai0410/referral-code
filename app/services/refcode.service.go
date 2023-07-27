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
	var refcode models.User
	code, _ := refcode.FindByAddress(address)
	if code == -1 {
		val, err := cache.Incr(constants.CacheCounter)
		if err != nil {
			log.Fatal("[SERVICE #2] " + err.Error())
		}
		code = val
		// SaveRefCode(&schemas.RefCodeRequest{
		// 	Address: address,
		// 	User: val,
		// })
		model := models.User{
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
	model := models.User{
		Address: address,
	}
	user, err := model.FindDocsByAddress(address)
	if err != nil {
		return nil, err
	}
	resp := new(schemas.TrackingResponse)
	resp.Count = user.Counter
	resp.Rate = user.Rate
	resp.Level = user.Level
	resp.WithdrawAvailable = user.WithdrawAvailable
	return resp, nil
}
