package internal

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
	"strconv"
)

type Configuration struct {
	DbHost          string
	DbPort          int
	DbName          string
	DbUser          string
	DbPass          string
	VapidPublicKey  string
	VapidPrivateKey string
}

var config Configuration

func GetConfig() Configuration {
	if config == (Configuration{}) {
		InitConfig()
	}

	return config
}

func InitConfig() {
	_ = godotenv.Load(fmt.Sprintf(".env%s", getEnvExtension()))

	dbPort, _ := strconv.Atoi(getEnv("DB_PORT", "27017"))

	config = Configuration{
		DbHost:          getEnv("DB_HOST", "localhost"),
		DbPort:          dbPort,
		DbName:          getEnv("DB_NAME", "candy-fight"),
		DbUser:          getEnv("DB_USER", "candy-fight"),
		DbPass:          getEnv("DB_PASS", "candy-fight"),
		VapidPublicKey:  getEnv("VAPID_PUBLIC_KEY", ""),
		VapidPrivateKey: getEnv("VAPID_PRIVATE_KEY", ""),
	}
}

func getEnvExtension() string {
	extension := ""
	if env := os.Getenv("APP_ENV"); env == "prod" {
		extension = ".prod"
	}

	return extension
}

func getEnv(key string, defaultValue string) string {
	v, set := os.LookupEnv(key)

	if set {
		return v
	}
	return defaultValue
}
