package main

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/antonholmquist/jason"
)

// Hub maintains the set of active clients and broadcasts messages to the
// clients.
type Hub struct {
	// Game which clients interact with
	game *Game

	// Registered clients.
	clients map[*Client]bool

	// Inbound messages from the clients.
	broadcast chan []byte

	// Register requests from the clients.
	register chan *Client

	// Unregister requests from clients.
	unregister chan *Client
}

func newHub(g *Game) *Hub {
	return &Hub{
		game:       g,
		broadcast:  make(chan []byte),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		clients:    make(map[*Client]bool),
	}
}

func (h *Hub) run() {
	// continuously update clients of new game state
	go func() {
		for {
			gameStateJSON, _ := json.Marshal(&h.game.GameState)

			for client := range h.clients {
				client.send <- gameStateJSON
			}

			// artificial delay
			time.Sleep(1 * time.Millisecond)
		}
	}()

	// continuously listen for clients and input
	go func() {
		for {
			select {
			case client := <-h.register:
				h.clients[client] = true
			case client := <-h.unregister:
				if _, ok := h.clients[client]; ok {
					// remove player from game
					h.game.RemovePlayers <- client.id

					// remove client from hub
					delete(h.clients, client)
					close(client.send)
				}
			case message := <-h.broadcast:
				// extract expected values from message
				json, _ := jason.NewObjectFromBytes([]byte(message))
				id, _ := json.GetInt64("id")
				x, _ := json.GetInt64("newTargetPosition", "x")
				y, _ := json.GetInt64("newTargetPosition", "y")

				// print read values
				fmt.Printf("PlayerInput: {id: %d, {x: %d, y: %d}}\n", id, x, y)

				// create PlayerInput struct
				input := PlayerInput{
					Id:                int(id),
					NewTargetPosition: Location{X: int(x), Y: int(y)},
				}

				// send player input to game
				h.game.PlayerInputs <- input
			}
		}
	}()
}
