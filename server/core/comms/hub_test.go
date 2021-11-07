package comms_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/atulanand206/find/server/core/actions"
	"github.com/atulanand206/find/server/core/comms"
	"github.com/atulanand206/find/server/core/models"
	gonanoid "github.com/matoous/go-nanoid/v2"
	"github.com/stretchr/testify/assert"
)

func TestHub(t *testing.T) {
	t.Run("New hub", func(t *testing.T) {
		hub := comms.NewHub()
		assert.NotNil(t, hub, "Hub should not be nil")
		assert.Equal(t, 0, len(hub.Clients), "Hub should have 0 clients")
	})

	t.Run("New hub with 1 client", func(t *testing.T) {
		hub := RunningHub(t)
		client := comms.NewClient(hub, nil)
		assert.NotNil(t, client, "Client should not be nil")
		hub.Register <- client
		time.Sleep(10 * time.Millisecond)
		assert.Equal(t, 1, len(hub.Clients), "Hub should have 1 client")
	})

	t.Run("New hub with 4 clients", func(t *testing.T) {
		hub := RunningHub(t)
		NewClient(t, hub)
		NewClient(t, hub)
		NewClient(t, hub)
		NewClient(t, hub)
		time.Sleep(10 * time.Millisecond)
		assert.Equal(t, 4, len(hub.Clients), "Hub should have 4 clients")
	})

	t.Run("Broadcast Hello to a hub with 4 clients", func(t *testing.T) {
		hub := RunningHubWithClients(t, 4)
		assert.NotNil(t, hub, "Hub should not be nil")
		msg := models.WebsocketMessageCreator{}.InitWebSocketMessage(actions.BEGIN, "Hello")
		targets := make(map[string]bool)
		for client := range hub.Clients {
			targets[client.PlayerId] = true
		}
		hub.Broadcast(msg, targets)
		for client := range hub.Clients {
			select {
			case message, _ := <-client.Send:
				assert.Equal(t, []byte(comms.SerializeMessage(msg)), message, "Message must be same as broadcasted")
			}
		}
	})
}

func RunningHubWithClients(t *testing.T, n int) *comms.Hub {
	hub := RunningHub(t)
	for i := 0; i < n; i++ {
		NewClient(t, hub)
	}
	time.Sleep(10 * time.Millisecond)
	assert.Equal(t, n, len(hub.Clients), fmt.Sprintf("Hub should have %d clients", n))
	return hub
}

func RunningHub(t *testing.T) *comms.Hub {
	hub := comms.NewHub()
	assert.NotNil(t, hub, "Hub should not be nil")
	go hub.Run()
	return hub
}

func NewClient(t *testing.T, hub *comms.Hub) *comms.Client {
	client := comms.NewClient(hub, nil)
	testPlayerId, _ := gonanoid.New(10)
	client.SetPlayerId(testPlayerId)
	assert.NotNil(t, client, "Client should not be nil")
	hub.Register <- client
	return client
}
