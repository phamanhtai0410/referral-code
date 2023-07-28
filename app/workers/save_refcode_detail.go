package workers

import (
	"encoding/json"
	"log"
	"time"

	"example.com/refcode/v1/app/models"
	"example.com/refcode/v1/app/schemas"
	"example.com/refcode/v1/platform/cache"
	"github.com/streadway/amqp"

	"example.com/refcode/v1/pkg/constants"
)

func SaveUser(msg *amqp.Delivery) {
	if msg.MessageId == constants.MsgSaveUserDetail {
		req := new(schemas.RefCodeRequest)
		err := json.Unmarshal(msg.Body, req)

		if err != nil {
			log.Fatal("[WORKER #3] failed to decode refcode: ", err)
		}
		id, _ := cache.Incr(constants.CacheCounter)
		model := models.User{
			ReferralCode:      req.Domain,
			Id:                id,
			Address:           req.Address,
			Created:           time.Now().UTC(),
			Counter:           0,
			Rate:              0,
			WithdrawAvailable: 0,
			Level:             "",
			LastUpdated:       time.Now().UTC(),
		}
		exits := model.IsExits(req.Address, req.Domain)
		if exits {
			return
		}
		err = model.Save()
		if err != nil {
			log.Fatal(constants.LogWorkerSaveCodeDetails, err)
		}
		log.Println(constants.LogWorkerSaveCodeDetails+" INSERTED OK: ", model)
		// msg.Ack(true)
	}

}
