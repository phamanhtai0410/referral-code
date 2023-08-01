package tasks

import (
	"time"

	"example.com/refcode/v1/app/models"
	"example.com/refcode/v1/pkg/constants"
	"github.com/RichardKnop/machinery/v1/log"
)

func SaveRefCodeUsed(referralCode, domain, address string, price float64) error {
	model := models.CodeUsed{
		ReferralCode: referralCode,
		Domain:       domain,
		Price:        price,
		Address:      address,
		Created:      time.Now().UTC(),
	}
	log.DEBUG.Println(constants.LogWorkerSaveCodeUsed+" INSERTED OK: ", model)
	return model.Save()
}
