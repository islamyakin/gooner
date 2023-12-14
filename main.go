package main

import (
	"gonef/server"
)

func main() {

	router := server.SetupRouter()
	port := ":9090"
	server.StartServer(port, router)
}
