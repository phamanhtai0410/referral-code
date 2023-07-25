package workers

import (
	"encoding/json"
	"example.com/refcode/v1/app/models"
	"example.com/refcode/v1/app/schemas"
	"example.com/refcode/v1/pkg/constants"
	"example.com/refcode/v1/platform/cache"
	"fmt"
	"github.com/streadway/amqp"
	"log"
	"sync"
	"time"
)

var lock sync.Mutex

func SaveRefCodeUsed(msg *amqp.Delivery) {
	if msg.MessageId == constants.MsgSaveRefCodeUsed {
		req := new(schemas.RefCodeUsedRequest)
		err := json.Unmarshal(msg.Body, req)
		if err != nil {
			log.Fatal("failed to decode refcode info: ", err)
		}
		model := models.RefCodeUsed{
			RefCode: req.RefCode,
			Domain:  req.Domain,
			Price:   req.Price,
			Address: req.Address,
			Created: time.Now().UTC().Format("2006-01-02T15:04:05.000Z"),
		}
		err = model.Save()
		if err != nil {
			log.Fatal(constants.LogWorkerSaveCodeUsed, err)
		}
		log.Println(constants.LogWorkerSaveCodeUsed+" INSERTED OK: ", model)
		lock.Lock()
		_, err = cache.Incr(fmt.Sprintf("CODE #%s", req.RefCode))
		if err != nil {
			log.Fatal("[WORKER] Save RefCode ", err)
			return
		}
		lock.Unlock()
	}
}
