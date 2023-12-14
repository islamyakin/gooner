package internal

import (
	"bufio"
	"github.com/docker/distribution/context"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strings"
)

type LogEntry struct {
	Message string `json:"message"`
}

func GetContainerLogs(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	containerId := vars["id"]

	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		log.Println("Error creating Docker client:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	options := types.ContainerLogsOptions{
		ShowStdout: true,
		ShowStderr: true,
		Follow:     false,
	}

	out, err := cli.ContainerLogs(context.Background(), containerId, options)
	if err != nil {
		log.Println("Error getting container logs:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer out.Close()

	// Salin log dari output ke response
	scanner := bufio.NewScanner(out)
	for scanner.Scan() {
		line := scanner.Text()
		// Jangan kirim karakter null
		line = strings.ReplaceAll(line, "\x00", "")
		// Tulis log ke response
		_, err := w.Write([]byte(line + "\n"))
		if err != nil {
			log.Println("Error writing log response:", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	if err := scanner.Err(); err != nil {
		log.Println("Error reading log:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
