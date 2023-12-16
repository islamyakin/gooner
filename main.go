package main

import (
	"github.com/islamyakin/gooner/server"
)

func main() {

	router := server.SetupRouter()
	port := ":9090"
	server.StartServer(port, router)
}
