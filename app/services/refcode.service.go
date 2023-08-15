package services

import (
	"log"

	"example.com/refcode/v1/app/models"
	"example.com/refcode/v1/app/schemas"
	"example.com/refcode/v1/app/tasks"
	"example.com/refcode/v1/pkg/utils"
	"example.com/refcode/v1/pkg/workers"
)

func SaveRefCodeInfo(req *schemas.RefCodeUsedRequest) {
	err := workers.Delay(
		"Worker.SaveCodeUsed",
		tasks.SaveRefCodeUsed,
		req.ReferralCode,
		req.Domain,
		req.Address,
		req.TransactionHash,
		req.Price,
	)
	if err != nil {
		log.Fatal(err)
	}
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
		Pending:      utils.Floor(user.Pending),
	}, nil
}

func WithdrawHistory(domain string) ([]models.WithdrawHistory, error) {
	var model models.WithdrawHistory
	return model.GetDocumentByDomain(domain)
}
