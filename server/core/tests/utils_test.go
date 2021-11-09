package tests_test

import (
	"os"
	"testing"

	"github.com/atulanand206/find/server/core/db"
	"github.com/atulanand206/find/server/core/tests"
	"github.com/atulanand206/go-mongo"
	"github.com/stretchr/testify/assert"
)

func TestSetup(t *testing.T) {
	t.Run("should setup the database", func(t *testing.T) {
		teardown := tests.Setup(t)
		defer teardown(t)
		assert.Equal(t, "mongodb://localhost:27017", mongo.ClientUrl)
		assert.Equal(t, "binquiz-test", db.Database)
	})

	t.Run("should setup the environment variables", func(t *testing.T) {
		teardown := tests.Setup(t)
		defer teardown(t)
		assert.Equal(t, "aedsaddsad05442b3fe16f2e72d1d497bf14ea9cb", os.Getenv("CLIENT_SECRET"))
		assert.Equal(t, "ae05442b3fe16f2e72d1d497bf14ea9dsafdsadsacb", os.Getenv("REFRESH_CLIENT_SECRET"))
		assert.Equal(t, "2", os.Getenv("TOKEN_EXPIRE_MINUTES"))
		assert.Equal(t, "10", os.Getenv("REFRESH_TOKEN_EXPIRE_MINUTES"))
	})
}

func TestRandomString(t *testing.T) {
	t.Run("should create a player", func(t *testing.T) {
		randomString := tests.TestRandomString()
		assert.NotEmpty(t, randomString, "random string should not be empty")
	})
}

func TestPlayer(t *testing.T) {
	t.Run("should create a player", func(t *testing.T) {
		player := tests.TestPlayer()
		assert.NotEmpty(t, player.Id, "Player id should not be empty")
		assert.NotEmpty(t, player.Name, "Player name should not be empty")
		assert.NotEmpty(t, player.Email, "Player email should not be empty")
	})
}

func TestTeams(t *testing.T) {
	t.Run("should create a team", func(t *testing.T) {
		quizId := tests.TestRandomString()
		team := tests.TestTeam(quizId)
		assert.NotEmpty(t, team.Id, "Team id should not be empty")
		assert.Equal(t, quizId, team.QuizId, "Team quiz id should be equal to the quiz id")
		assert.NotEmpty(t, team.Name, "Team name should not be empty")
		assert.Zero(t, team.Score, "Team score should be zero")
	})

	t.Run("should create 4 teams", func(t *testing.T) {
		quizId := tests.TestRandomString()
		teams := tests.TestTeams(quizId, 4)
		for _, team := range teams {
			assert.NotEmpty(t, team.Id, "Team id should not be empty")
			assert.Equal(t, quizId, team.QuizId, "Team quiz id should be equal to the quiz id")
			assert.NotEmpty(t, team.Name, "Team name should not be empty")
			assert.Zero(t, team.Score, "Team score should be zero")
		}
	})
}

func TestGame(t *testing.T) {
	t.Run("should create a game spec", func(t *testing.T) {
		specs := tests.TestSpecs()
		assert.NotEmpty(t, specs.Name, "quiz name should not be empty")
		assert.Equal(t, 2, specs.Teams, "quiz teams should be equal to 2")
		assert.Equal(t, 2, specs.Players, "quiz players should be equal to 2")
		assert.Equal(t, 10, specs.Questions, "quiz questions should be equal to 10")
		assert.Equal(t, 2, specs.Rounds, "quiz rounds should be equal to 2")
		assert.Equal(t, 10, specs.Points, "quiz points should be equal to 10")
	})

	t.Run("should create a game", func(t *testing.T) {
		quiz := tests.TestGame()
		assert.NotEmpty(t, quiz.Id, "quiz id should not be empty")
		assert.NotEmpty(t, quiz.QuizMaster.Id, "quiz master id should not be empty")
		assert.NotEmpty(t, quiz.QuizMaster.Name, "quiz master name should not be empty")
		assert.NotEmpty(t, quiz.QuizMaster.Email, "quiz master email should not be empty")
		assert.NotEmpty(t, quiz.Specs.Name, "quiz name should not be empty")
		assert.Equal(t, 2, quiz.Specs.Teams, "quiz teams should be equal to 2")
		assert.Equal(t, 2, quiz.Specs.Players, "quiz players should be equal to 2")
		assert.Equal(t, 10, quiz.Specs.Questions, "quiz questions should be equal to 10")
		assert.Equal(t, 2, quiz.Specs.Rounds, "quiz rounds should be equal to 2")
		assert.Equal(t, 16, quiz.Specs.Points, "quiz points should be equal to 10")
		assert.True(t, quiz.Active, "quiz must be active")
		assert.Empty(t, quiz.Tags, "quiz tags must be empty")
	})
}
