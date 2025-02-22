package main

import (
	"discord_dota2/internal/api/router"
	discordInit "discord_dota2/internal/discord/init"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	err := discordInit.InitDiscordClient()
	if err != nil {
		panic(err)
	}
	mainRouter := mux.NewRouter()
	router.SetIntegrationRouter(mainRouter)
	log.Println("application started")
	err = http.ListenAndServe(":52424", mainRouter)
	if err != nil {
		panic(err)
	}
}
