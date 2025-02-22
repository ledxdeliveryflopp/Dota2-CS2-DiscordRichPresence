package main

import (
	"discord_dota2/internal/api/router"
	discordInit "discord_dota2/internal/discord/init"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	mainRouter := mux.NewRouter()
	router.SetIntegrationRouter(mainRouter)
	err := discordInit.InitDiscordClient()
	if err != nil {
		panic(err)
	}
	err = http.ListenAndServe(":52424", mainRouter)
	log.Println("application started")
	if err != nil {
		panic(err)
	}
}
