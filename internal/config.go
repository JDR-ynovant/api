package internal

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
	"strconv"
)

type Configuration struct {
	// DB
	DbHost string
	DbPort int
	DbName string
	DbUser string
	DbPass string
	// Webpush
	VapidPublicKey  string
	VapidPrivateKey string
	// Notification strings
	StringsFilePath string
	// Game Rules
	RuleAttackRange   int
	RuleMoveRange     int
	RuleBaseDamage    int
	RuleBloodSugarCap int
	RuleWalkableQuota uint
	RuleObstacleQuota uint
	// Cors policy
	CorsAllowOrigins string
	CorsAllowHeaders string
	CorsAllowMethods string
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

	moveRange, _ := strconv.Atoi(getEnv("CANDY_FIGHT_RULE_MOVE_RANGE", "2"))
	attackRange, _ := strconv.Atoi(getEnv("CANDY_FIGHT_RULE_ATTACK_RANGE", "1"))
	baseDamage, _ := strconv.Atoi(getEnv("CANDY_FIGHT_RULE_BASE_DAMAGE", "1"))
	maxLife, _ := strconv.Atoi(getEnv("CANDY_FIGHT_RULE_BLOOD_SUGAR_CAP", "10"))
	walkableQuota, _ := strconv.ParseUint(getEnv("CANDY_FIGHT_RULE_WALKABLE_QUOTA", "11"), 10, 64)
	obstacleQuota, _ := strconv.ParseUint(getEnv("CANDY_FIGHT_RULE_OBSTACLE_QUOTA", "2"), 10, 64)

	config = Configuration{
		DbHost:            getEnv("DB_HOST", "localhost"),
		DbPort:            dbPort,
		DbName:            getEnv("DB_NAME", "candy-fight"),
		DbUser:            getEnv("DB_USER", "candy-fight"),
		DbPass:            getEnv("DB_PASS", "candy-fight"),
		VapidPublicKey:    getEnv("VAPID_PUBLIC_KEY", ""),
		VapidPrivateKey:   getEnv("VAPID_PRIVATE_KEY", ""),
		StringsFilePath:   getEnv("STRINGS_FILE_PATH", "strings.yml"),
		RuleAttackRange:   attackRange,
		RuleMoveRange:     moveRange,
		RuleBaseDamage:    baseDamage,
		RuleBloodSugarCap: maxLife,
		RuleWalkableQuota: uint(walkableQuota),
		RuleObstacleQuota: uint(obstacleQuota),
		CorsAllowOrigins:  getEnv("CORS_ALLOW_ORIGINS", "*"),
		CorsAllowHeaders:  getEnv("CORS_ALLOW_HEADERS", ""),
		CorsAllowMethods:  getEnv("CORS_ALLOW_METHODS", "GET,POST,HEAD,PUT,DELETE,OPTIONS"),
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
