package configs

import (
	"encoding/json"
	"os"

	log "github.com/sirupsen/logrus"
)

type SteamSettings struct {
	SteamID string `json:"steam_id"`
}

func (s *SteamSettings) InitSettings() {
	fileData, err := os.ReadFile("settings.json")
	if err != nil {
		log.Panicln(err)
	}
	err = json.Unmarshal(fileData, &s)
	if err != nil {
		log.Panicln(err)
	}
	return
}
