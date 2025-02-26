package main

import (
	"discord_dota2_cs2/internal/api/router"
	discordInit "discord_dota2_cs2/internal/discord/init"
	"fmt"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
)

func initLogrus() {
	file, err := os.OpenFile("logrus.log", os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Panicln(err)
		return
	}
	log.SetOutput(file)
	log.Info("Logrus initer")
}

func main() {
	initLogrus()
	discordInit.InitDiscordClient()
	go discordInit.CheckGameIsRunning()
	mainRouter := mux.NewRouter()
	router.SetIntegrationRouter(mainRouter)
	fmt.Println("start api")
	err := http.ListenAndServe(":52424", mainRouter)
	if err != nil {
		log.Panicln("error while starting api:", err)
	}
}
