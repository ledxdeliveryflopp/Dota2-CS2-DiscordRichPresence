package router

import (
	"discord_dota2/internal/api/handler"
	"github.com/gorilla/mux"
)

func SetIntegrationRouter(router *mux.Router) {
	router.HandleFunc("/", handler.HandleGameStateResponse).Methods("POST")
}
