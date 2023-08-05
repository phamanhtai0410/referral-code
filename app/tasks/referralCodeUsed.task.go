package tasks

import (
	"fmt"
	"time"

	"example.com/refcode/v1/app/models"
	"example.com/refcode/v1/pkg/constants"
	"example.com/refcode/v1/platform/sentry"
	"github.com/RichardKnop/machinery/v1/log"
)

func SaveRefCodeUsed(referralCode, domain, address, transactionHash string, price float64) error {
	model := models.CodeUsed{
		ReferralCode:    referralCode,
		Domain:          domain,
		Price:           price,
		Address:         address,
		Created:         time.Now().UTC(),
		TransactionHash: transactionHash,
	}
	log.DEBUG.Println(constants.LogWorkerSaveCodeUsed+" INSERTED OK: ", model)
	sentry.CaptureMessage(fmt.Sprintf("%s INSERTED OK: %s-%s-%s-%f", constants.LogWorkerSaveCodeUsed, model.Address, model.Domain, model.ReferralCode, model.Price))
	return model.Save()
}
