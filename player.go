package main

import (
	"math/rand"
)

// Player a player of the game
type Player struct {
	ID int
	// TODO restrict to X or O
	MoveType string
}

// NewPlayer make a new player
func NewPlayer(moveType string) Player {
	id := rand.Intn(100000)

	player := Player{id, moveType}

	return player
}
