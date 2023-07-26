package queue

import (
	"log"
	"sync"

	"example.com/refcode/v1/pkg/configs"
	"example.com/refcode/v1/pkg/constants"

	"github.com/streadway/amqp"
)

var conn *amqp.Connection
var channel *amqp.Channel
var queue *amqp.Queue
var lock sync.Mutex

func init() {
	var err error
	conn, err = amqp.Dial(configs.BrokerUrl)
	if err != nil {
		log.Fatal("[RabbitMQ] Cannot connect to RabbitMQ Server: " + err.Error())
	}
	log.Println("[RabbitMQ] connection successfully")
	channel, err = conn.Channel()
	if err != nil {
		log.Fatal("Cannot create channel to RabbitMQ Server: " + err.Error())
	}
	queue, err = NewQueue(constants.WorkerQueue)
	if err != nil {
		log.Fatal("Cannot create queue to RabbitMQ Server: " + err.Error())
	}
	log.Println("[RabbitMQ] channel successfully")
}

func Shutdown() {
	log.Println("[RabbitMQ] Shutting down RabbitMQ connection ...")
	if channel != nil {
		if err := channel.Close(); err != nil {
			log.Fatal(err)
		}
		// log.Println("[RabbitMQ] channel successfully closed")
	}
	if conn != nil {
		if err := conn.Close(); err != nil {
			log.Fatal(err)
		}
		// log.Println("[RabbitMQ] connection successfully closed")
	}
}

func NewQueue(queueName string) (*amqp.Queue, error) {
	if queue == nil {
		lock.Lock()
		if queue == nil {
			q, err := channel.QueueDeclare(
				queueName,
				true,
				false,
				false,
				false,
				nil,
			)
			queue = &q
			lock.Unlock()
			return &q, err
		}
		lock.Unlock()
	}
	return queue, nil
}

func Publish(queueName string, msgID string, msg []byte) error {
	err := channel.Publish(
		"",
		queue.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        msg,
			MessageId:   msgID,
		},
	)
	return err
}

func Consume(queueName string) (<-chan amqp.Delivery, error) {
	tag := "worker-refcode"
	msg, err := channel.Consume(
		queue.Name,
		tag,
		true,
		false,
		false,
		false,
		nil,
	)
	return msg, err
}
