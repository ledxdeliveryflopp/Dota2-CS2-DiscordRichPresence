package handler

import (
	csgoTypes "discord_dota2_cs2/internal/api/csgo_types"
	dotaTypes "discord_dota2_cs2/internal/api/dota_types"
	discordHandler "discord_dota2_cs2/internal/discord/handler"
	"fmt"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func HandleDotaGameStateResponse(writer http.ResponseWriter, request *http.Request) {
	var dotaResponse dotaTypes.GameDotaResponse
	err := dotaResponse.DecodeGameInfo(request.Body)
	if err != nil {
		log.Error("error while encoding json: ", err)
	}
	log.Infof("Dota response: %+v", dotaResponse)
	log.Infof("Dota response: %+v", dotaResponse)
	defer request.Body.Close()
	errorChan := make(chan error)
	successChan := make(chan bool)
	go discordHandler.SetDotaPresence(successChan, errorChan, &dotaResponse)
	select {
	case err = <-errorChan:
		log.Error("error while set dota presence: ", err)
		fmt.Println("error while set dota presence, see log file for info")
		close(errorChan)
		close(successChan)
		return
	case <-successChan:
		close(errorChan)
		close(successChan)
		return
	}
}

func HandleCsGoGameStateResponse(writer http.ResponseWriter, request *http.Request) {
	var csgoResponse csgoTypes.GameCsgoResponse
	err := csgoResponse.DecodeGameInfo(request.Body)
	if err != nil {
		log.Error("error while encoding json: ", err)
	}
	log.Infof("Cs Go response: %+v", csgoResponse)
	defer request.Body.Close()
	errorChan := make(chan error)
	successChan := make(chan bool)
	go discordHandler.SetCsGoPresence(successChan, errorChan, &csgoResponse)
	select {
	case err = <-errorChan:
		log.Error("error while set csgo presence: ", err)
		fmt.Println("error while set csgo presence, see log file for info")
		close(errorChan)
		close(successChan)
		return
	case <-successChan:
		close(errorChan)
		close(successChan)
		return
	}
}
