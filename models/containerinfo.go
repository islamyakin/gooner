package models

type ContainerInfo struct {
	ContainerID string   `json:"container_id"`
	Name        string   `json:"name_container"`
	State       string   `json:"state"`
	Status      string   `json:"status"`
	Ports       []string `json:"ports,omitempty"`
}
