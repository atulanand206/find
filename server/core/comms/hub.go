package comms

import (
	"bytes"
	"encoding/json"

	"github.com/atulanand206/find/server/core/models"
	"github.com/atulanand206/find/server/core/services"
)

// Hub maintains the set of active clients and broadcasts messages to the
// clients.
type Hub struct {
	// Registered Clients.
	Clients map[*Client]bool

	// Inbound messages from the clients.
	broadcast chan Message

	// Register requests from the clients.
	Register chan *Client

	// Unregister requests from clients.
	Unregister chan *Client

	Controller services.Service
}

type Message struct {
	targets map[string]bool

	msg []byte
}

func NewHub(controller services.Service) *Hub {
	return &Hub{
		broadcast:  make(chan Message),
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Clients:    make(map[*Client]bool),
		Controller: controller,
	}
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.Register:
			h.Clients[client] = true
		case client := <-h.Unregister:
			if _, ok := h.Clients[client]; ok {
				msg, targets, _ := h.Controller.DeletePlayerLiveSession(client.PlayerId)
				delete(h.Clients, client)
				close(client.Send)
				go h.Broadcast(msg, targets)
			}
		case message := <-h.broadcast:
			for client := range h.Clients {
				if message.targets[client.PlayerId] {
					select {
					case client.Send <- message.msg:
					default:
						close(client.Send)
						delete(h.Clients, client)
					}
				}
			}
		}
	}
}

func (hub *Hub) HandleMessages(input []byte, fromClient *Client) (err error) {
	request, err := models.DecodeWebSocketRequest(input)
	if err != nil {
		return
	}
	response, targets, err := hub.Handle(request, fromClient)
	if err != nil {
		return
	}
	hub.Broadcast(response, targets)
	return
}

func (hub *Hub) Broadcast(response models.WebsocketMessage, targets map[string]bool) {
	message := SerializeMessage(response)
	hub.broadcast <- Message{msg: message, targets: targets}
}

func SerializeMessage(response interface{}) []byte {
	output, err := json.Marshal(response)
	if err != nil {
		return nil
	}
	message := bytes.TrimSpace(bytes.Replace(output, newline, space, -1))
	return message
}
