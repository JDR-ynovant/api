package internal

import (
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

type Strings struct {
	NotificationPlayerIsDead string `yaml:"notification_player_is_dead"`
	NotificationPlayerTurn   string `yaml:"notification_player_turn"`
	NotificationPlayerWin    string `yaml:"notification_player_win"`
	NotificationGameStart    string `yaml:"notification_game_start"`
}

var strings Strings

func GetStrings() Strings {
	if strings == (Strings{}) {
		InitStrings()
	}

	return strings
}

func InitStrings() {
	config := GetConfig()
	f, err := os.Open(config.StringsFilePath)
	if err != nil {
		log.Fatal(err.Error())
	}

	defer f.Close()

	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&strings)
	if err != nil {
		log.Fatal(err.Error())
	}

	log.Printf("strings red from `%s` file.\n", config.StringsFilePath)
}
