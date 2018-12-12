package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/pkg/errors"
)

// NOTE: capitalisation denoted visibility outside current package

// Game is a struct representing a noughts and crosses game
type Game struct {
	ID          int        `json:"id"`
	PlayerTurn  int        `json:"playerTurn"`
	Board       [][]string `json:"board"`
	PlayerOneID int        `json:"player1Id"`
	PlayerTwoID int        `json:"player2Id"`
}

func (g Game) checkWinCondition() {

}

func (g *Game) makeMove(playerID int, x int, y int) error {
	// TODO join the game if free space, otherwise return an error
	// if playerID != g.Player1ID && playerID != g.Player2ID {
	// 	return errors.New("blah")
	// }

	// TODO For some reason this doesn't seem to be working
	if playerID != g.PlayerTurn {
		return errors.New("It is not your turn")
	}

	if x > 2 || y > 2 {
		return errors.New("Error: space does not exist")
	}

	// TODO try and catch index errors
	if g.Board[y][x] != "" {
		return errors.New("Error: space is not empty")
	}

	g.Board[y][x] = "X"

	if g.PlayerTurn == g.PlayerOneID {
		g.PlayerTurn = g.PlayerTwoID
	} else {
		g.PlayerTurn = g.PlayerOneID
	}

	return nil
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
