package configs

import (
	"log"
	"os"
	"strconv"

	_ "github.com/joho/godotenv/autoload" // load .env file automatically
)

var (
	StageStatus string
	_           string
	_           int
	_           int

	_ string
	_ int
	_ string
	_ int

	DbType string
	DbHost string
	DbPort int
	_      string
	_      string
	DbName string
	_      string
	_      int
	_      int
	_      int

	_        string
	_        int
	_        string
	_        int
	CacheUrl string

	BrokerUrl string

	BlockChainRPC string
	ChainID       int
	ContractAddr  string
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
	_ = os.Getenv("SERVER_HOST")
	_ = convertEnvToInt("SERVER_PORT")
	_ = convertEnvToInt("SERVER_READ_TIMEOUT")
	_ = os.Getenv("JWT_SECRET_KEY")
	_ = convertEnvToInt("JWT_SECRET_KEY_EXPIRE_MINUTES_COUNT")
	_ = os.Getenv("JWT_REFRESH_KEY")
	_ = convertEnvToInt("JWT_REFRESH_KEY_EXPIRE_HOURS_COUNT")
	DbType = os.Getenv("DB_TYPE")
	DbHost = os.Getenv("DB_HOST")
	DbPort = convertEnvToInt("DB_PORT")
	_ = os.Getenv("DB_USER")
	_ = os.Getenv("DB_PASSWORD")
	DbName = os.Getenv("DB_NAME")
	_ = os.Getenv("DB_SSL_MODE")
	_ = convertEnvToInt("DB_MAX_CONNECTIONS")
	_ = convertEnvToInt("DB_MAX_IDLE_CONNECTIONS")
	_ = convertEnvToInt("DB_MAX_LIFETIME_CONNECTIONS")
	_ = os.Getenv("REDIS_HOST")
	_ = convertEnvToInt("REDIS_PORT")
	_ = os.Getenv("REDIS_PASSWORD")
	_ = convertEnvToInt("REDIS_DB_NUMBER")
	BrokerUrl = os.Getenv("BROKER_URL")
	CacheUrl = os.Getenv("REDIS_URL")
	BlockChainRPC = os.Getenv("BLOCKCHAIN_RPC")
	ChainID = convertEnvToInt("CHAIN_ID")
	ContractAddr = os.Getenv("CONTRACT_ADDR")
}

type TransactOptsConfig struct {
	PrivateKey string `json:"private_key"`
	ChainID    int64  `json:"chain_id"`
}
