
package internal

type Configuration struct {
	DbHost string
	DbPort int
	DbName string
}

var config Configuration

func GetConfig() Configuration {
	if config == (Configuration{}) {
		InitConfig()
	}

	return config
}

func InitConfig() {
	config = Configuration{
		DbHost: "localhost",
		DbPort: 27017,
		DbName: "candy-fight",
	}
}