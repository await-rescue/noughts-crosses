package main

import (
	"fmt"
	"math/rand"

	"github.com/pkg/errors"
)

// Game is a struct representing a noughts and crosses game
type Game struct {
	ID          int        `json:"id"`
	PlayerTurn  int        `json:"playerTurn"`
	Board       [][]string `json:"board"`
	PlayerOneID int        `json:"playerOneId"`
	PlayerTwoID int        `json:"playeTwoId"`
	Status      string     `json:"status"`
}

func (g *Game) checkWinCondition(player *Player) {
	// check rows
	for i := 0; i < len(g.Board); i++ {
		if g.Board[i][0] == player.MoveType && g.Board[i][1] == player.MoveType && g.Board[i][2] == player.MoveType {
			g.Status = fmt.Sprintf("%s wins", player.Name)
			return
		}
	}

	// check cols
	for i := 0; i < len(g.Board); i++ {
		if g.Board[0][i] == player.MoveType && g.Board[1][i] == player.MoveType && g.Board[2][i] == player.MoveType {
			g.Status = fmt.Sprintf("%s wins", player.Name)
			return
		}
	}

	// check diagonals
	if g.Board[0][0] == player.MoveType && g.Board[1][1] == player.MoveType && g.Board[2][2] == player.MoveType {
		g.Status = fmt.Sprintf("%s wins", player.Name)
		return
	}

	if g.Board[2][2] == player.MoveType && g.Board[1][1] == player.MoveType && g.Board[0][0] == player.MoveType {
		g.Status = fmt.Sprintf("%s wins", player.Name)
		return
	}

	// If we find any empty spaces, do nothing and return, otherwise game is a draw
	for _, row := range g.Board {
		for _, square := range row {
			if square == "" {
				return
			}
		}
	}

	g.Status = "Draw"
}

func (g *Game) makeMove(playerID int, x int, y int) error {

	if g.Status != "active" {
		return errors.New(fmt.Sprintf("Game is over.\nStatus: %s", g.Status))
	}

	if playerID != g.PlayerTurn {
		return errors.New("It is not your turn")
	}

	if x > 2 || y > 2 {
		return errors.New("Error: space does not exist")
	}

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

	g.checkWinCondition(player)

	return nil
}

// NewGame creates a new game with a random id
func NewGame(playerOne *Player, playerTwo *Player) Game {
	// Ideally we'd enforce uniqueness and store game in a DB
	id := rand.Intn(100000)
	rows := [][]string{{"", "", ""}, {"", "", ""}, {"", "", ""}}

	game := Game{id, playerOne.ID, rows, playerOne.ID, playerTwo.ID, "active"}

	return game
}
