package server

import (
	"github.com/gorilla/mux"
	"gonef/internal"
	"log"
	"net/http"
)

func SetupRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/api/running-containers", internal.ListRunningContainers).Methods("GET")
	r.HandleFunc("/api/container/{id}/start", internal.StartContainer).Methods("POST")
	r.HandleFunc("/api/container/{id}/stop", internal.StopContainer).Methods("POST")
	r.HandleFunc("/api/container/{id}/restart", internal.RestartContainer).Methods("POST")
	return r

}

func StartServer(port string, router *mux.Router) {
	log.Printf("Server is running on port %s...\n", port)
	log.Fatal(http.ListenAndServe(port, router))
}
