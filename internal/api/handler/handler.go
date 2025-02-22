package handler

import (
	"discord_dota2/internal/api/types"
	"discord_dota2/internal/discord/handler"
	"log"
	"net/http"
)

func HandleGameStateResponse(writer http.ResponseWriter, request *http.Request) {
	var dotaResponse types.GameResponse
	err := dotaResponse.DecodeGameInfo(request.Body)
	log.Println(dotaResponse)
	if err != nil {
		log.Fatalln("error while encoding json: ", err)
	}
	defer request.Body.Close()
	errorChan := make(chan error)
	successChan := make(chan bool)

	go handler.SetDiscordPresence(successChan, errorChan, &dotaResponse.Player, &dotaResponse.Hero)
	select {
	case err = <-errorChan:
		log.Println("error while set presence: ", err)
		close(errorChan)
		close(successChan)
		return
	case <-successChan:
		close(errorChan)
		close(successChan)
		return
	}
}
