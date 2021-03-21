package main

import (
	"encoding/json"
	"fmt"
)

// Hub maintains the set of active clients and broadcasts messages to the
// clients.
type Hub struct {
	// Registered clients.
	clients map[*Client]bool

	// Inbound messages from the clients.
	broadcast chan []byte

	// Register requests from the clients.
	register chan *Client

	// Unregister requests from clients.
	unregister chan *Client
}

func newHub() *Hub {
	return &Hub{
		broadcast:  make(chan []byte),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		clients:    make(map[*Client]bool),
	}
}

func (h *Hub) run() {
	for {
		select {
		case client := <-h.register:
			h.clients[client] = true
		case client := <-h.unregister:
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				close(client.send)
			}
		case message := <-h.broadcast:
			fmt.Printf("%s\n", message)

			// extract name value from message
			var bodyJSON map[string]interface{}
			json.Unmarshal([]byte(message), &bodyJSON)

			/*input := PlayerInput{
				Id:                bodyJSON["id"].(int),
				NewTargetPosition: Location{X: bodyJSON["targetPosition"]["x"].(int), Y: bodyJSON["targetPosition"]["y"].(int)},
			}*/

			locationJSON := bodyJSON["targetPosition"].(map[string]interface{})
			fmt.Printf("x: %f\n", locationJSON["x"].(float64))

			fmt.Printf("id: %d\n", int(bodyJSON["id"].(float64)))
		}
	}
}
