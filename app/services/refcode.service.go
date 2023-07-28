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

func ReferralCodeHandle(req *schemas.RefCodeRequest) (int64, error) {

	var user models.User
	address, id := user.OwnerOf(req.Domain)
	if address != "" && address != req.Address {
		return -1, errors.New("domain already exists")
	}
	if id == -1 {
		if user.WalletExists(req.Address) {
			return -1, errors.New("address already exists")
		}
		_id, err := cache.Incr(constants.CacheCounter)
		if err != nil {
			return -1, err
		}
		req.Id = _id
		data, err := json.Marshal(req)
		if err != nil {
			return -1, err
		}
		if err = queue.Publish(
			constants.WorkerQueue,
			constants.MsgSaveUserDetail,
			data,
		); err != nil {
			return -1, err
		}
		log.Println("[SERVICES #9] Publishing to queue ...")
		return _id, nil
	}
	return id, nil
}

func SaveRefCodeInfo(req *schemas.RefCodeUsedRequest) error {
	var user models.User
	addr, id := user.OwnerOf(req.ReferralCode)
	if addr == "" {
		return errors.New("referral code does not exist")
	}
	req.Id = id
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
	// resp := new(schemas.TrackingResponse)
	// resp.Count = user.Counter
	// resp.Rate = user.Rate
	// resp.Level = user.Level
	// resp.WithdrawAvailable = user.WithdrawAvailable
	// resp.ReferralCode = user.ReferralCode
	// resp.Id = user.Id
	// return resp, nil
	return &schemas.TrackingResponse{
		Count:             user.Counter,
		Rate:              user.Rate,
		Level:             user.Level,
		WithdrawAvailable: user.WithdrawAvailable,
		ReferralCode:      user.ReferralCode,
	}, nil
}
