package configs

import (
	"github.com/RichardKnop/machinery/v1/config"
)

type Worker struct {
	Task   map[string]interface{}
	Config *config.Config
}

func WorkerBaseSetting(queueName, broker, resultBackend string) *config.Config {
	return &config.Config{
		Broker:          broker,
		DefaultQueue:    queueName,
		ResultBackend:   resultBackend,
		ResultsExpireIn: 3600,
		AMQP: &config.AMQPConfig{
			Exchange:      "machinery_exchange",
			ExchangeType:  "direct",
			BindingKey:    queueName,
			PrefetchCount: 3,
		},
	}
}
