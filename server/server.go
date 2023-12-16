package server

import (
	"github.com/gorilla/mux"
	"github.com/islamyakin/gooner/middleware"
	"github.com/islamyakin/gooner/internal"
	"log"
	"net/http"
)

func SetupRouter() *mux.Router {
	r := mux.NewRouter()

	r.Use(middleware.LoggerMiddleware)

	r.HandleFunc("/", internal.LandingGooner).Methods("GET")


	allowedIPs := []string{"192.168.1.1"}
	ipRestrictedRouter := r.PathPrefix("/api").Subrouter()
	ipRestrictedRouter.Use(middleware.IPWhitelistMiddleware(allowedIPs, nil))
	ipRestrictedRouter.HandleFunc("/api/running-containers", internal.ListRunningContainers).Methods("GET")
	ipRestrictedRouter.HandleFunc("/api/container/{id}/start", internal.StartContainer).Methods("POST")
	ipRestrictedRouter.HandleFunc("/api/container/{id}/stop", internal.StopContainer).Methods("POST")
	ipRestrictedRouter.HandleFunc("/api/container/{id}/restart", internal.RestartContainer).Methods("POST")
	ipRestrictedRouter.HandleFunc("/api/container/{id}/logs", internal.GetContainerLogs).Methods("GET")
	return r

}

func StartServer(port string, router *mux.Router) {
	log.Printf("Server is running on port %s...\n", port)
	log.Fatal(http.ListenAndServe(port, router))
}
