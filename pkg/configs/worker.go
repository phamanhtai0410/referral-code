package configs

import (
	"github.com/RichardKnop/machinery/v1/config"
)


type Worker struct {
	Task   map[string]interface{}
	Config *config.Config
}
