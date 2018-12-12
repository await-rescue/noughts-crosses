package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var games = make(map[int]*Game)
var players = make(map[int]*Player)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/xo/create/", CreateGame)
	router.HandleFunc("/xo/make-move/", PlayerMove)

	log.Fatal(http.ListenAndServe(":8080", router))
}
