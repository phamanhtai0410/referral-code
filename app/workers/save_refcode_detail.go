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
			log.Fatal("failed to decode refcode: ", err)
		}
		model := models.RefCode{
			RefCode: req.RefCode,
			Address: req.Address,
			Created: time.Now().UTC().Format("2006-01-02T15:04:05.000Z"),
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
