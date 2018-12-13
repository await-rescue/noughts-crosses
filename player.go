package main

import (
	"math/rand"
)

// Player a player of the game
type Player struct {
	ID       int
	Name     string
	MoveType string
}

// NewPlayer make a new player
func NewPlayer(name string, moveType string) Player {
	id := rand.Intn(200000)

	player := Player{id, name, moveType}

	return player
}
