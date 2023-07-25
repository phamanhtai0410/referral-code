package configs

import (
	"log"
	"os"
	"strconv"

	// _ "github.com/joho/godotenv/autoload" // load .env file automatically
)

var (
	StageStatus         string
	SERVER_HOST         string
	SERVER_PORT         int
	SERVER_READ_TIMEOUT int

	// JWT settings:
	JWT_SECRET_KEY                      string
	JWT_SECRET_KEY_EXPIRE_MINUTES_COUNT int
	JWT_REFRESH_KEY                     string
	JWT_REFRESH_KEY_EXPIRE_HOURS_COUNT  int

	// Database settings:
	DB_TYPE                     string
	DB_HOST                     string
	DB_PORT                     int
	DB_USER                     string
	DB_PASSWORD                 string
	DB_NAME                     string
	DB_SSL_MODE                 string
	DB_MAX_CONNECTIONS          int
	DB_MAX_IDLE_CONNECTIONS     int
	DB_MAX_LIFETIME_CONNECTIONS int

	// Redis settings:
	REDIS_HOST      string
	REDIS_PORT      int
	REDIS_PASSWORD  string
	REDIS_DB_NUMBER int
	REDIS_URL       string

	// Message queue:
	BROKER_URL string
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
	SERVER_HOST = os.Getenv("SERVER_HOST")
	SERVER_PORT = convertEnvToInt("SERVER_PORT")
	SERVER_READ_TIMEOUT = convertEnvToInt("SERVER_READ_TIMEOUT")
	JWT_SECRET_KEY = os.Getenv("JWT_SECRET_KEY")
	JWT_SECRET_KEY_EXPIRE_MINUTES_COUNT = convertEnvToInt("JWT_SECRET_KEY_EXPIRE_MINUTES_COUNT")
	JWT_REFRESH_KEY = os.Getenv("JWT_REFRESH_KEY")
	JWT_REFRESH_KEY_EXPIRE_HOURS_COUNT = convertEnvToInt("JWT_REFRESH_KEY_EXPIRE_HOURS_COUNT")
	DB_TYPE = os.Getenv("DB_TYPE")
	DB_HOST = os.Getenv("DB_HOST")
	DB_PORT = convertEnvToInt("DB_PORT")
	DB_USER = os.Getenv("DB_USER")
	DB_PASSWORD = os.Getenv("DB_PASSWORD")
	DB_NAME = os.Getenv("DB_NAME")
	DB_SSL_MODE = os.Getenv("DB_SSL_MODE")
	DB_MAX_CONNECTIONS = convertEnvToInt("DB_MAX_CONNECTIONS")
	DB_MAX_IDLE_CONNECTIONS = convertEnvToInt("DB_MAX_IDLE_CONNECTIONS")
	DB_MAX_LIFETIME_CONNECTIONS = convertEnvToInt("DB_MAX_LIFETIME_CONNECTIONS")
	REDIS_HOST = os.Getenv("REDIS_HOST")
	REDIS_PORT = convertEnvToInt("REDIS_PORT")
	REDIS_PASSWORD = os.Getenv("REDIS_PASSWORD")
	REDIS_DB_NUMBER = convertEnvToInt("REDIS_DB_NUMBER")
	BROKER_URL = os.Getenv("BROKER_URL")
	REDIS_URL = os.Getenv("REDIS_URL")
}
