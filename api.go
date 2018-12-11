package main

import (
	"encoding/json"
	"net/http"
)

// Move serialises json data
type Move struct {
	GameID   int `json:"gameID"`
	PlayerID int `json:"playerID"`
	X        int `json:"x"`
	Y        int `json:"y"`
}

// Error for returning request errors
type Error struct {
	Error string `json:"error"`
}

// CreateGame creates a struct Game http GET :8080/xo/create/
func CreateGame(rw http.ResponseWriter, r *http.Request) {
	game := NewGame()
	games[game.ID] = game

	data, err := json.Marshal(game)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	rw.WriteHeader(200)
	rw.Header().Set("Content-Type", "application/json")
	rw.Write(data)
}

// PlayerMove adds either an X or O to the board
func PlayerMove(rw http.ResponseWriter, r *http.Request) {
	move := Move{}

	err := json.NewDecoder(r.Body).Decode(&move)
	if err != nil {
		panic(err)
	}
	game := games[move.GameID]
	game.makeMove(move.PlayerID, move.X, move.Y)

	data, err := json.Marshal(game)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	rw.WriteHeader(200)
	rw.Header().Set("Content-Type", "application/json")
	rw.Write(data)
}
