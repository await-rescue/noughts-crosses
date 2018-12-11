package main

import (
	"encoding/json"
	"net/http"
)

// StringResponse returns a basic string response
type StringResponse struct {
	Result string
}

// Ping tests the API is up
func Ping(rw http.ResponseWriter, r *http.Request) {
	ping := StringResponse{"Pong"}

	data, err := json.Marshal(ping)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	rw.WriteHeader(200)
	rw.Header().Set("Content-Type", "application/json")
	rw.Write(data)
}

// CreateGame creates a struct Game http GET :8080/xo/create/
func CreateGame(rw http.ResponseWriter, r *http.Request) {
	game := NewGame()
	// TODO can we store the games in a global map (or sqlite)?

	data, err := json.Marshal(game)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	rw.WriteHeader(200)
	rw.Header().Set("Content-Type", "application/json")
	rw.Write(data)
}
