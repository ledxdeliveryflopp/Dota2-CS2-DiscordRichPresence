package handler

import (
	csgoTypes "discord_dota2_cs2/internal/api/csgo_types"
	dotaTypes "discord_dota2_cs2/internal/api/dota_types"
	config "discord_dota2_cs2/internal/configs"
	discordHandler "discord_dota2_cs2/internal/discord/handler"
	"discord_dota2_cs2/internal/discord/types"
	"fmt"
	"net/http"
)

func HandleDotaGameStateResponse(writer http.ResponseWriter, request *http.Request) {
	var dotaResponse dotaTypes.GameDotaResponse
	err := dotaResponse.DecodeGameInfo(request.Body)
	if err != nil {
		config.DotaLog.Println("-------------------------------------------")
		config.DotaLog.Error("error while encoding json: ", err)
		return
	}
	defer request.Body.Close()
	errorChan := make(chan types.DotaPresenceError)
	successChan := make(chan bool)
	go discordHandler.SetDotaPresence(successChan, errorChan, &dotaResponse)
	select {
	case err := <-errorChan:
		config.DotaLog.Println("-------------------------------------------")
		config.DotaLog.Error("error while set dota presence: ", err.Err)
		config.DotaLog.Infof("State: %s", err.Presence.State)
		config.DotaLog.Infof("Details: %s", err.Presence.Details)
		config.DotaLog.Infof("LargeImage: %s", err.Presence.MainImage)
		config.DotaLog.Infof("LargeText: %s", err.Presence.HeroReadableName)
		config.DotaLog.Infof("SmallImage: %s", err.Presence.SmallImage)
		config.DotaLog.Infof("SmallText: %s", "Dota 2")
		fmt.Println("error while set dota presence, see dota.log file for info")
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
		config.CsGoLog.Println("-------------------------------------------")
		config.CsGoLog.Error("error while encoding json: ", err)
		return
	}
	defer request.Body.Close()
	errorChan := make(chan types.CsGoPresenceError)
	successChan := make(chan bool)
	go discordHandler.SetCsGoPresence(successChan, errorChan, &csgoResponse)
	select {
	case err := <-errorChan:
		config.CsGoLog.Println("-------------------------------------------")
		config.CsGoLog.Error("Error while set CS:GO presence: ", err)
		config.CsGoLog.Infof("State: %s", err.Presence.State)
		config.CsGoLog.Infof("Details: %s", err.Presence.Details)
		config.CsGoLog.Infof("LargeImage: %s", "main")
		config.CsGoLog.Infof("LargeText: %s", "CS:GO 2")
		config.CsGoLog.Infof("SmallImage: %s", "team image")
		config.CsGoLog.Infof("SmallText: %s", "team name")
		fmt.Println("error while set csgo presence, see csgo.log file for info")
		close(errorChan)
		close(successChan)
		return
	case <-successChan:
		close(errorChan)
		close(successChan)
		return
	}
}
