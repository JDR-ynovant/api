
package internal

type Configuration struct {
	DbHost string
	DbPort int
	DbName string
}

var config Configuration

func GetConfig() Configuration {
	return config
}

func InitConfig() {
	config = Configuration{}
}