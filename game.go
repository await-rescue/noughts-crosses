package main

import (
	"fmt"
	"math/rand"

	"github.com/pkg/errors"
)

// NOTE: capitalisation denoted visibility outside current package

// Game is a struct representing a noughts and crosses game
// TODO might want to just ref a player directly
type Game struct {
	ID          int        `json:"id"`
	PlayerTurn  int        `json:"playerTurn"`
	Board       [][]string `json:"board"`
	PlayerOneID int        `json:"playerOneId"`
	PlayerTwoID int        `json:"playeTwoId"`
}

func (g Game) checkWinCondition() {

}

func (g *Game) makeMove(playerID int, x int, y int) error {

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

	player := players[playerID]

	g.Board[y][x] = player.MoveType

	if g.PlayerTurn == g.PlayerOneID {
		g.PlayerTurn = g.PlayerTwoID
	} else {
		g.PlayerTurn = g.PlayerOneID
	}

	return nil
}

// NewGame creates a new game with a random id
func NewGame(playerOne *Player, playerTwo *Player) Game {
	rows := [][]string{{"", "", ""}, {"", "", ""}, {"", "", ""}}

	// Ideally we'd enforce uniqueness and store game in a DB
	id := rand.Intn(100000)

	game := Game{id, playerOne.ID, rows, playerOne.ID, playerTwo.ID}
	str := fmt.Sprintf("Created game with id %d", id)
	fmt.Println(str)
	return game
}
