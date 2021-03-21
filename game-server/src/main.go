package main

import (
//"encoding/json"
//"fmt"
//"io/ioutil"
//"log"
//"net/http"
//"time"
)

const ()

func main() {
	// create and run a Game instance
	game := newGame()
	go game.run()

	// create and run a Hub instance
	hub := newHub()
	go hub.run()

	/*
		newPlayer := Player{
			Id:              69,
			DisplayName:     "Lukas",
			LastUpdateTime:  int(time.Now().UnixNano() / 1000000),
			CurrentPosition: Location{X: 0, Y: 0},
			TargetPosition:  Location{X: 0, Y: 0},
			MoveSpeed:       100,
		}
		game.GameState.Players = append(game.GameState.Players, newPlayer)

		fmt.Printf("%+v\n", game.GameState)

		input := PlayerInput{
			Id: 69,
			NewTargetPosition: Location{
				X: 20,
				Y: 90,
			},
		}
		game.PlayerInputs <- input

		time.Sleep(500 * time.Millisecond)
		fmt.Printf("%+v\n", game.GameState)
		time.Sleep(500 * time.Millisecond)
		fmt.Printf("%+v\n", game.GameState)
	*/

	/*hub := newHub()
	go hub.run()
	http.HandleFunc("/player-create", playerCreatePost)
	http.HandleFunc("/player-input-ws", func(w http.ResponseWriter, r *http.Request) { serveWs(hub, w, r) })
	log.Fatal("ListenAndServe", http.ListenAndServe(":80", nil))
	*/
}

/*
func playerCreatePost(w http.ResponseWriter, r *http.Request) {
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
	for i := 0; i < len(players); i++ {
		if players[i].Id > max {
			max = players[i].Id
		}

		if players[i].Id == id {
			id = max + 1
		}
	}

	// create new player and append to players
	newPlayer := Player{
		Id:              id,
		Name:            name,
		LastUpdateTime:  int(time.Now().UnixNano() / 1000000),
		CurrentPosition: Location{X: 100, Y: 100},
		TargetPosition:  Location{X: 100, Y: 100},
		MoveSpeed:       100,
	}
	players = append(players, newPlayer)

	// respond to request with JSON of new player
	newPlayerJSON, err := json.Marshal(&newPlayer)
	log.Println(players)
	w.Header().Set("Access-Control-Allow-Origin", "http://subparprogramming.cf")
	w.Write([]byte(newPlayerJSON))
}
*/
