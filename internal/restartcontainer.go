package internal

import (
	"fmt"
	"github.com/docker/distribution/context"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func RestartContainer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	containerID := vars["id"]

	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		log.Println("Error creating Docker client:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := cli.ContainerRestart(context.Background(), containerID, container.StopOptions{}); err != nil {
		log.Println("Error restarting container:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_, err = fmt.Fprintf(w, "Container %s restarted\n", containerID)
	if err != nil {
		log.Println("Error writing response:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
