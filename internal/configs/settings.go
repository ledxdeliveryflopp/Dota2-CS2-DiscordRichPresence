package configs

import (
	"encoding/json"
	"errors"
	"os"

	log "github.com/sirupsen/logrus"
)

type SteamSettings struct {
	SteamID string `json:"steam_id"`
}

var Settings *SteamSettings

func InitSteamSettings() {
	fileData, err := os.ReadFile("settings.json")
	if err != nil {
		log.Panicln(err)
	}
	var settings SteamSettings
	err = json.Unmarshal(fileData, &settings)
	if err != nil {
		log.Panicln(err)
	}
	if settings == (SteamSettings{}) {
		err = errors.New("empty settings file")
		log.Panicln(err)
	}
}
