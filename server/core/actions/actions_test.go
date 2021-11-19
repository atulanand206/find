package actions_test

import (
	"testing"

	"github.com/atulanand206/find/server/core/actions"
	"github.com/stretchr/testify/assert"
)

func TestAction(t *testing.T) {
	t.Run("should convert to string", func(t *testing.T) {
		assert.Equal(t, "BEGIN", actions.BEGIN.String())
		assert.Equal(t, "SPECS", actions.SPECS.String())
		assert.Equal(t, "JOIN", actions.JOIN.String())
		assert.Equal(t, "WATCH", actions.WATCH.String())
		assert.Equal(t, "CREATE", actions.CREATE.String())
		assert.Equal(t, "START", actions.START.String())
		assert.Equal(t, "HINT", actions.HINT.String())
		assert.Equal(t, "RIGHT", actions.RIGHT.String())
		assert.Equal(t, "PASS", actions.PASS.String())
		assert.Equal(t, "NEXT", actions.NEXT.String())
		assert.Equal(t, "REVEAL", actions.REVEAL.String())
		assert.Equal(t, "DROP", actions.DROP.String())
		assert.Equal(t, "SCORE", actions.SCORE.String())
		assert.Equal(t, "REFRESH", actions.REFRESH.String())
		assert.Equal(t, "S_REFRESH", actions.S_REFRESH.String())
		assert.Equal(t, "FINISH", actions.FINISH.String())
		assert.Equal(t, "extend", actions.Extend.String())
		assert.Equal(t, "ACTIVE", actions.ACTIVE.String())
		assert.Equal(t, "S_ACTIVE", actions.S_ACTIVE.String())
		assert.Equal(t, "S_GAME", actions.S_GAME.String())
		assert.Equal(t, "S_JOIN", actions.S_JOIN.String())
		assert.Equal(t, "S_PLAYER", actions.S_PLAYER.String())
		assert.Equal(t, "S_START", actions.S_START.String())
		assert.Equal(t, "S_HINT", actions.S_HINT.String())
		assert.Equal(t, "S_RIGHT", actions.S_RIGHT.String())
		assert.Equal(t, "S_PASS", actions.S_PASS.String())
		assert.Equal(t, "S_NEXT", actions.S_NEXT.String())
		assert.Equal(t, "S_SCORE", actions.S_SCORE.String())
		assert.Equal(t, "s_question", actions.S_Question.String())
		assert.Equal(t, "s_answer", actions.S_Answer.String())
		assert.Equal(t, "S_OVER", actions.S_OVER.String())
		assert.Equal(t, "failure", actions.Failure.String())
		assert.Equal(t, "EMPTY", actions.EMPTY.String())
	})
}

func TestRole(t *testing.T) {
	t.Run("should convert to string", func(t *testing.T) {
		assert.Equal(t, "QUIZMASTER", actions.QUIZMASTER.String())
		assert.Equal(t, "PLAYER", actions.PLAYER.String())
		assert.Equal(t, "AUDIENCE", actions.AUDIENCE.String())
		assert.Equal(t, "TEAM", actions.TEAM.String())
	})
}
