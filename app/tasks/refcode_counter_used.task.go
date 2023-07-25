package tasks

import (
	"example.com/refcode/v1/app/models"
	"example.com/refcode/v1/pkg/constants"
	"github.com/streadway/amqp"
)

func RefCodeCounterUsed(msg *amqp.Delivery) {
	if msg.MessageId == constants.MsgRefCodeCounter {
		code := string(msg.Body)
		model := models.RefCode{
			RefCode: code,
		}
		count, isExist := model.IsExits(code)
		if !isExist {
			return
		}
		count++
		// TODO:
	}
}
