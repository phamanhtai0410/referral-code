package services

import (
	"encoding/json"
	"errors"
	"log"

	"example.com/refcode/v1/app/models"
	"example.com/refcode/v1/platform/cache"

	"example.com/refcode/v1/app/schemas"
	"example.com/refcode/v1/pkg/constants"
	"example.com/refcode/v1/platform/queue"
)

func ReferralCodeHandle(req *schemas.RefCodeRequest) error {

	var user models.User
	address := user.OwnerOf(req.Domain)
	if address != "" && address != req.Address {
		return errors.New("domain already exists")
	}

	data, err := json.Marshal(req)
	if err != nil {
		return err
	}
	if err = queue.Publish(
		constants.WorkerQueue,
		constants.MsgSaveUserDetail,
		data,
	); err != nil {
		return err
	}
	log.Println("[SERVICES #9] Publishing to queue ...")
	return nil
}

func SaveRefCodeInfo(req *schemas.RefCodeUsedRequest) error {
	currentCode, err := cache.Get(constants.CacheCounter)
	if err != nil {
		return err
	}
	if req.Id <= constants.CacheCounterBegin || req.Id > currentCode {
		return errors.New("referral code does not exist yet")
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
		constants.MsgSaveUserDetail,
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
	resp.ReferralCode = user.ReferralCode
	resp.Id = user.Id
	return resp, nil
}
