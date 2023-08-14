package configs

import (
	"log"
	"os"
	"strconv"

	_ "github.com/joho/godotenv/autoload" // load .env file automatically
)

var (
	StageStatus string

	DbType        string
	DbHost        string
	DbPort        int
	DbName        string
	ResultBackend string

	BrokerUrl string

	BlockChainRPC string
	ChainID       int
	ContractAddr  string
	SentryDNS     string
)

func convertEnvToInt(key string) int {
	valueStr := os.Getenv(key)
	valueInt, err := strconv.Atoi(valueStr)
	if err != nil {
		log.Fatal("ERROR: "+key+": "+valueStr+" ", err)
	}
	return valueInt
}

func init() {
	StageStatus = os.Getenv("STAGE_STATUS")
	DbType = os.Getenv("DB_TYPE")
	DbHost = os.Getenv("DB_HOST")
	DbPort = convertEnvToInt("DB_PORT")
	DbName = os.Getenv("DB_NAME")
	BrokerUrl = os.Getenv("BROKER_URL")
	ResultBackend = os.Getenv("RESULT_BACKEND")
	BlockChainRPC = os.Getenv("BLOCKCHAIN_RPC")
	ChainID = convertEnvToInt("CHAIN_ID")
	ContractAddr = os.Getenv("CONTRACT_ADDR")
	SentryDNS = os.Getenv("SENTRY_DNS")
}

type TransactOptsConfig struct {
	PrivateKey string `json:"private_key"`
	ChainID    int64  `json:"chain_id"`
}
