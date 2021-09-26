package core

import (
	"io"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

type (
	WebsocketMessage struct {
		Person Player `json:"person"`
	}
)

var clients = make(map[*websocket.Conn]bool)
var broadcaster = make(chan WebsocketMessage)
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func handleMessages() {
	for {
		// grab any next message from channel
		msg := <-broadcaster
		messageClients(msg)
	}
}

func messageClients(msg WebsocketMessage) {
	for client := range clients {
		err := client.WriteJSON(msg)
		if err != nil && unsafeError(err) {
			log.Printf("error: %v", err)
			client.Close()
			delete(clients, client)
		}
	}
}

// If a message is sent while a client is closing, ignore the error
func unsafeError(err error) bool {
	return !websocket.IsCloseError(err, websocket.CloseGoingAway) && err != io.EOF
}

func HandlerWebSockets(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}
	// ensure connection close when function returns
	defer ws.Close()
	clients[ws] = true

	for {
		var msg WebsocketMessage
		// Read in a new message as JSON and map it to a Message object
		err := ws.ReadJSON(&msg)
		if err != nil {
			delete(clients, ws)
			break
		}
		// send new message to the channel
		broadcaster <- msg
	}
}
