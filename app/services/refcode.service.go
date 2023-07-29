package services

import (
	"encoding/json"
	"errors"
	"log"

	"example.com/refcode/v1/app/models"
	"example.com/refcode/v1/app/schemas"
	"example.com/refcode/v1/pkg/constants"
	"example.com/refcode/v1/platform/queue"
)

func SaveRefCodeInfo(req *schemas.RefCodeUsedRequest) error {
	data, err := json.Marshal(req)
	if err != nil {
		return errors.New("[SERVICE #3] failed to encode " + err.Error())
	}
	if err = queue.Publish(
		constants.WorkerQueue,
		constants.MsgSaveRefCodeUsed,
		data,
	); err != nil {
		return errors.New("[SERVICE #4] Failed to publish msg " + err.Error())
	}
	log.Println("[SERVICES #5] Publishing to queue ...")
	return nil
}

func RefCodeTracking(referralCode string) (*schemas.TrackingResponse, error) {
	model := models.User{
		ReferralCode: referralCode,
	}
	user, err := model.FindDocsByReferralCode(referralCode)
	if err != nil {
		return nil, err
	}
	return &schemas.TrackingResponse{
		Count:        user.Counter,
		Rate:         user.Rate,
		Level:        user.Level,
		TotalEarn:    user.TotalEarn,
		ReferralCode: user.ReferralCode,
	}, nil
}

func WithdrawHistory(domain string) ([]models.WithdrawHistory, error) {
	var model models.WithdrawHistory
	return model.GetDocumentByDomain(domain)
}
