package main

import (
	"math/rand"
	"time"
)

// Player a player of the game
type Player struct {
	ID int
	// TODO restrict to X or O
	MoveType string
}

// NewPlayer make a new player
func NewPlayer(moveType string) Player {
	rand.Seed(time.Now().UnixNano())
	id := rand.Intn(100000)

	player := Player{id, moveType}

	return player
}
