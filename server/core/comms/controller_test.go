package comms_test

import (
	"testing"

	"github.com/atulanand206/find/server/core/actions"
	"github.com/atulanand206/find/server/core/tests"
	"github.com/stretchr/testify/assert"
)

func TestHandleMessages(t *testing.T) {
	t.Run("Handle a string as web socket message", func(t *testing.T) {
		hub := tests.RunningHubWithClients(t, 2)
		msg := tests.TestMessage(actions.BEGIN, "Hello")
		_, _, err := hub.Handle(msg, tests.ClientX(hub))
		assert.NotNil(t, err, "error must be present")
	})
}
