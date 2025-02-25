package router

import (
	"discord_dota2_cs2/internal/api/handler"
	"github.com/gorilla/mux"
)

func SetIntegrationRouter(router *mux.Router) {
	gsiRouter := router.PathPrefix("/handler").Subrouter()
	gsiRouter.HandleFunc("/dota/", handler.HandleDotaGameStateResponse).Methods("POST")
	gsiRouter.HandleFunc("/cs/", handler.HandleCsGoGameStateResponse).Methods("POST")
}
