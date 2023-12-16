package internal

import (
	"encoding/json"
	"fmt"
	"github.com/docker/distribution/context"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/client"
	"github.com/islamyakin/gooner/models"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func ListRunningContainers(w http.ResponseWriter, _ *http.Request) {
	cli, err := client.NewClientWithOpts(client.FromEnv)

	if err != nil {
		log.Println("Error creating Docker client:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	newFilter := filters.NewArgs()
	newFilter.Add("status", "running")

	containers, err := cli.ContainerList(context.Background(), types.ContainerListOptions{
		All:     false,
		Filters: newFilter,
	})
	if err != nil {
		log.Println("Error listing containers:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// slice untuk menyimpan informasi lisContainer
	containerInfoList := make([]models.ContainerInfo, 0)

	// Mengisi informasi lisContainer
	for _, lisContainer := range containers {
		containerInfo := models.ContainerInfo{
			ContainerID: lisContainer.ID,
			Name:        strings.TrimPrefix(lisContainer.Names[0], "/"),
			State:       lisContainer.State,
			Status:      lisContainer.Status,
		}

		// Mendapatkan informasi port lisContainer
		for _, port := range lisContainer.Ports {
			containerInfo.Ports = append(containerInfo.Ports, fmt.Sprintf("%s:%s->%s/%s", port.IP, strconv.Itoa(int(port.PublicPort)), strconv.Itoa(int(port.PrivatePort)), port.Type))
		}

		containerInfoList = append(containerInfoList, containerInfo)
	}

	// Konversi slice  menjadi JSON
	jsonData, err := json.Marshal(containerInfoList)
	if err != nil {
		log.Println("Error marshalling JSON:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(jsonData)
	if err != nil {
		log.Println("Error writing JSON:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
