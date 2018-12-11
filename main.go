package main

import (
    "log"
    "net/http"

    "github.com/gorilla/mux"
)

func main() {
    router := mux.NewRouter()

    router.HandleFunc("/", Ping)
    router.HandleFunc("/xo/create/", CreateGame)

    log.Fatal(http.ListenAndServe(":8080", router))
}
