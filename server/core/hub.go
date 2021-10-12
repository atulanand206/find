package core

import "fmt"

// Hub maintains the set of active clients and broadcasts messages to the
// clients.
type Hub struct {
	// Registered clients.
	clients map[*Client]bool

	// Inbound messages from the clients.
	broadcast chan Message

	// Register requests from the clients.
	register chan *Client

	// Unregister requests from clients.
	unregister chan *Client
}

type Message struct {
	targets map[string]bool

	msg []byte
}

func NewHub() *Hub {
	return &Hub{
		broadcast:  make(chan Message),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		clients:    make(map[*Client]bool),
	}
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.register:
			fmt.Println(client)
			h.clients[client] = true
		case client := <-h.unregister:
			if _, ok := h.clients[client]; ok {
				msg, targets, _ := Controller.playerService.DeletePlayerLiveSession(client.playerId)
				delete(h.clients, client)
				close(client.send)
				go h.Broadcast(msg, targets)
			}
		case message := <-h.broadcast:
			for client := range h.clients {
				if message.targets[client.playerId] {
					select {
					case client.send <- message.msg:
					default:
						close(client.send)
						delete(h.clients, client)
					}
				}
			}
		}
	}
}
