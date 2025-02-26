package configs

import (
	"encoding/json"
	"os"

	log "github.com/sirupsen/logrus"
)

type SteamSettings struct {
	SteamID string `json:"steamid"`
}

func (s *SteamSettings) InitSettings() {
	fileData, err := os.ReadFile("settings.json")
	if err != nil {
		log.Panicln(err)
	}
	var settings SteamSettings
	err = json.Unmarshal(fileData, &settings)
	if err != nil {
		log.Panicln(err)
	}
}
