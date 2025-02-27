package main

import (
	"discord_dota2_cs2/internal/api/router"
	config "discord_dota2_cs2/internal/configs"
	discordInit "discord_dota2_cs2/internal/discord/init"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	config.InitLogrus()
	discordInit.InitDiscordClient()
	go discordInit.CheckGameIsRunning()
	config.MainLog.Println("start check game status loop")
	mainRouter := mux.NewRouter()
	router.SetIntegrationRouter(mainRouter)
	fmt.Println("start api")
	err := http.ListenAndServe(":52424", mainRouter)
	if err != nil {
		config.MainLog.Panicln("error while starting api:", err)
	}
}
