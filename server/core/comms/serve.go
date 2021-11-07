package comms

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

// serveWs handles websocket requests from the peer.
func ServeWebsocket(hub *Hub, w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	client := NewClient(hub, conn)
	ServeBinquiz(hub, client)
}

func NewClient(hub *Hub, conn *websocket.Conn) *Client {
	return &Client{hub: hub, conn: conn, Send: make(chan []byte, 256)}
}

func ServeBinquiz(hub *Hub, client *Client) {
	client.hub.Register <- client
	go client.WritePump()
	go client.ReadPump()
}
