package internal

import "net/http"

func LandingGooner(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("Gooner is running...\n"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
