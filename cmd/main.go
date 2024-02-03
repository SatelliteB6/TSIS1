package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/SatelliteB6/TSIS1/pkg/champions"
	"github.com/gorilla/mux"
)

func main() {
	log.Println("starting API server")
	router := mux.NewRouter()

	// Specify endpoints
	router.HandleFunc("/health-check", HealthCheck).Methods("GET")
	router.HandleFunc("/champions", Champions).Methods("GET")

	http.Handle("/", router)

	// Start and listen to requests
	http.ListenAndServe(":8080", router)
}

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	log.Println("entering health check endpoint")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("API is up and running"))
}

func Champions(w http.ResponseWriter, r *http.Request) {
	log.Println("entering champions endpoint")
	var response champions.Response
	championsList := champions.PrepareChampions()

	response.Champions = championsList

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
