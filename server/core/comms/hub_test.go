package comms_test

import (
	"testing"
	"time"

	"github.com/atulanand206/find/server/core/actions"
	"github.com/atulanand206/find/server/core/comms"
	"github.com/atulanand206/find/server/core/db"
	"github.com/atulanand206/find/server/core/services"
	"github.com/atulanand206/find/server/core/tests"
	"github.com/stretchr/testify/assert"
)

func TestHub(t *testing.T) {
	t.Run("New hub", func(t *testing.T) {
		hub := comms.NewHub(services.Init(db.NewMockDb(false)))
		assert.NotNil(t, hub, "Hub should not be nil")
		assert.Equal(t, 0, len(hub.Clients), "Hub should have 0 clients")
	})

	t.Run("New hub with 1 client", func(t *testing.T) {
		hub := tests.RunningHub(t)
		client := comms.NewClient(hub, nil)
		assert.NotNil(t, client, "Client should not be nil")
		hub.Register <- client
		time.Sleep(10 * time.Millisecond)
		assert.Equal(t, 1, len(hub.Clients), "Hub should have 1 client")
	})

	t.Run("New hub with 4 clients", func(t *testing.T) {
		hub := tests.RunningHub(t)
		for i := 0; i < 4; i++ {
			tests.NewClient(t, hub)
		}
		time.Sleep(10 * time.Millisecond)
		assert.Equal(t, 4, len(hub.Clients), "Hub should have 4 clients")
	})

	t.Run("Broadcast Hello to a hub with 4 clients", func(t *testing.T) {
		hub := tests.RunningHubWithClients(t, 4)
		assert.NotNil(t, hub, "Hub should not be nil")
		msg := tests.TestMessage(actions.BEGIN, "Hello")
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
