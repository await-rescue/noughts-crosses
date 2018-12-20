
Simple API for noughts and crosses.


**To run:**

`go get github.com/sinksinksink/noughts-crosses`
or clone this repo and run `go get` from the folder for dependencies

`go run .`

**Create a game:**

`GET localhost:8080/xo/create/`
(make a note of the player IDs)

**Make a move:**

`POST localhost:8080/xo/make-move/`
example data to send with request: `{"gameID": 24728, "playerID": 128162, "x":1, "y":1}`

The coordinates are counted from the top left being 0,0

**Get the game status:**

`POST localhost:8080/xo/status/`
example data to send with request: `{"gameID": 24728}`
