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

// CreateGame creates a struct Game http GET :8080/xo/create/
func CreateGame(rw http.ResponseWriter, r *http.Request) {
	player1 := NewPlayer("X")
	player2 := NewPlayer("O")
	game := NewGame(&player1, &player2)

	games[game.ID] = &game
	players[player1.ID] = &player1
	players[player2.ID] = &player2

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
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	game := games[move.GameID]

	err = game.makeMove(move.PlayerID, move.X, move.Y)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}

	data, err := json.Marshal(game)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	rw.WriteHeader(200)
	rw.Header().Set("Content-Type", "application/json")
	rw.Write(data)
}
