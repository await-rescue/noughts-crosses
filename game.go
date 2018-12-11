package main

import (
	"fmt"
	"math/rand"
	"time"
)

// NOTE: capitalisation denoted visibility outside current package

// Game is a struct representing a noughts and crosses game
type Game struct {
	ID         int        `json:"id"`
	PlayerTurn int        `json:"playerTurn"`
	Board      [][]string `json:"board"`
	Player1ID  int        `json:"player1Id"`
	Player2ID  int        `json:"player2Id"`
}

func (g Game) checkWinCondition() {

}

func (g Game) makeMove(playerID int, x int, y int) {
	// TODO join the game if free space, otherwise return an error
	if playerID != g.Player1ID && playerID != g.Player2ID {
		fmt.Println("blah")
	}

	if g.Board[y][x] != "" {
		// TODO return an error so we can return it through json
		fmt.Println("Error: space is not empty")
	} else {
		g.Board[y][x] = "X"
	}
}

// NewGame creates a new game with a random id
func NewGame() Game {
	rows := [][]string{{"", "", ""}, {"", "", ""}, {"", "", ""}}

	// Ideally we'd enforce uniqueness and store game in a DB
	rand.Seed(time.Now().UnixNano())
	id := rand.Intn(100000)

	// TODO we should also make a player

	game := Game{id, 1, rows, 1, 2}
	str := fmt.Sprintf("Created game with id %d", id)
	fmt.Println(str)
	return game
}
