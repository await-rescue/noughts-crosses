package main
import (
    "fmt"
    "math/rand"
    "time"
)

// NOTE: capitalisation denoted visibility outside current package

// Game is a struct representing a noughts and crosses game
type Game struct {
    // TODO use a sub struct for each row?
    ID int `json:"id"`
    Rows []string `json:"rows"`
}

func (g Game) checkWinCondition() {

}

func (g Game) makeMove() {

}

// NewGame creates a new game with a random id
func NewGame() Game {
    rows := make([]string, 3)

    // Ideally we'd enforce uniqueness and store game in a DB
    rand.Seed(time.Now().UnixNano())
    id := rand.Intn(100000)

    game := Game{id, rows}
    str := fmt.Sprintf("Created game with id %d", id)
    fmt.Println(str)
    return game
}
