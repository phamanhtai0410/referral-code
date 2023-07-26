package workers

import (
	"encoding/json"
	"log"
	"time"

	"example.com/refcode/v1/app/models"
	"example.com/refcode/v1/app/schemas"
	"github.com/streadway/amqp"

	"example.com/refcode/v1/pkg/constants"
)

func SaveRefCode(msg *amqp.Delivery) {
	if msg.MessageId == constants.MsgSaveRefCodeDetail {
		req := new(schemas.RefCodeRequest)
		err := json.Unmarshal(msg.Body, req)
		if err != nil {
			log.Fatal("[WORKER #3] failed to decode refcode: ", err)
		}
		model := models.RefCode{
			RefCode: req.RefCode,
			Address: req.Address,
			Created: time.Now().UTC(),
			Counter: 0,
		}
		err = model.Save()
		if err != nil {
			log.Fatal(constants.LogWorkerSaveCodeDetails, err)
		}
		log.Println(constants.LogWorkerSaveCodeDetails+" INSERTED OK: ", model)
		// msg.Ack(true)
	}

}
