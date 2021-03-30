package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func main() {
	// create and run a Game instance
	game := newGame()
	go game.run()

	// create and run a Hub instance
	hub := newHub(game)
	go hub.run()

	http.HandleFunc("/player-create", func(w http.ResponseWriter, r *http.Request) { playerCreatePost(game, w, r) })
	http.HandleFunc("/player-input-ws", func(w http.ResponseWriter, r *http.Request) { serveWs(hub, w, r) })
	log.Fatal("ListenAndServe", http.ListenAndServe(":80", nil))
}

func playerCreatePost(g *Game, w http.ResponseWriter, r *http.Request) {
	// convert request body to byte array
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Print("[ERROR] ioutil.ReadAll:", err)
		return
	}

	// extract name value from body
	var bodyJSON map[string]interface{}
	json.Unmarshal([]byte(body), &bodyJSON)
	name := bodyJSON["name"].(string)

	// choose unique id for new player
	id := 0
	max := id
	for i := 0; i < len(g.GameState.Players); i++ {
		if g.GameState.Players[i].Id > max {
			max = g.GameState.Players[i].Id
		}

		if g.GameState.Players[i].Id == id {
			id = max + 1
		}
	}

	// create new player and append to players
	newPlayer := Player{
		Id:              id,
		DisplayName:     name,
		LastUpdateTime:  int(time.Now().UnixNano() / 1000000),
		CurrentPosition: Location{X: 100, Y: 100},
		TargetPosition:  Location{X: 100, Y: 100},
		MoveSpeed:       200,
	}
	g.GameState.Players = append(g.GameState.Players, newPlayer)

	// respond to request with JSON of new player
	newPlayerJSON, _ := json.Marshal(&newPlayer)
	w.Header().Set("Access-Control-Allow-Origin", "http://subparprogramming.cf")
	w.Write([]byte(newPlayerJSON))
}
