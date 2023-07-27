package workers

import (
	"encoding/json"
	"log"
	"time"

	"example.com/refcode/v1/app/models"
	"example.com/refcode/v1/app/schemas"
	"example.com/refcode/v1/pkg/constants"
	"github.com/streadway/amqp"
)

func SaveRefCodeUsed(msg *amqp.Delivery) {
	if msg.MessageId == constants.MsgSaveRefCodeUsed {
		req := new(schemas.RefCodeUsedRequest)
		err := json.Unmarshal(msg.Body, req)
		if err != nil {
			log.Fatal("[WORKER #4] failed to decode refcode info: ", err)
		}
		model := models.CodeUsed{
			RefCode: req.RefCode,
			Domain:  req.Domain,
			Price:   req.Price,
			Address: req.Address,
			Created: time.Now().UTC(),
		}
		err = model.Save()
		if err != nil {
			log.Fatal(constants.LogWorkerSaveCodeUsed, err)
		}
		log.Println(constants.LogWorkerSaveCodeUsed+" INSERTED OK: ", model)
	}
}
