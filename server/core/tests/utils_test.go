package tests_test

import (
	"os"
	"testing"

	"github.com/atulanand206/find/server/core/actions"
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

func TestQuestions(t *testing.T) {
	t.Run("should create an index", func(t *testing.T) {
		index := tests.TestIndex()
		assert.NotEmpty(t, index.Id, "index id should not be empty")
		assert.NotEmpty(t, index.Tag, "index tag should not be empty")
	})

	t.Run("should create 8 indexes", func(t *testing.T) {
		indexes := tests.TestIndexes(8)
		for _, index := range indexes {
			assert.NotEmpty(t, index.Id, "index id should not be empty")
			assert.NotEmpty(t, index.Tag, "index tag should not be empty")
		}
	})

	t.Run("should create a new question", func(t *testing.T) {
		question := tests.TestNewQuestion()
		assert.NotEmpty(t, question.Statements, "question statements should not be empty")
		assert.NotEmpty(t, question.Answer, "question answer should not be empty")
	})

	t.Run("should create a new question", func(t *testing.T) {
		newQuestion := tests.TestNewQuestion()
		assert.NotEmpty(t, newQuestion.Statements, "question statements should not be empty")
		assert.NotEmpty(t, newQuestion.Answer, "question answer should not be empty")
		index := tests.TestIndex()
		assert.NotEmpty(t, index.Id, "index id should not be empty")
		assert.NotEmpty(t, index.Tag, "index tag should not be empty")
		question := tests.TestQuestion(index.Tag, newQuestion)
		assert.NotEmpty(t, question.Id, "question id should not be empty")
		assert.Equal(t, index.Tag, question.Tag, "question tag should be equal to the index tag")
		assert.Equal(t, newQuestion.Statements, question.Statements, "question statements should be equal to the new question statements")
	})

	t.Run("should create a new answer", func(t *testing.T) {
		newQuestion := tests.TestNewQuestion()
		assert.NotEmpty(t, newQuestion.Statements, "question statements should not be empty")
		assert.NotEmpty(t, newQuestion.Answer, "question answer should not be empty")
		index := tests.TestIndex()
		assert.NotEmpty(t, index.Id, "index id should not be empty")
		assert.NotEmpty(t, index.Tag, "index tag should not be empty")
		question := tests.TestQuestion(index.Tag, newQuestion)
		assert.NotEmpty(t, question.Id, "question id should not be empty")
		assert.Equal(t, index.Tag, question.Tag, "question tag should be equal to the index tag")
		assert.Equal(t, newQuestion.Statements, question.Statements, "question statements should be equal to the new question statements")
		answer := tests.TestAnswer(question.Id, newQuestion.Answer)
		assert.NotEmpty(t, answer.Id, "answer id should not be empty")
		assert.Equal(t, question.Id, answer.QuestionId, "answer question id should be equal to the question id")
		assert.Equal(t, newQuestion.Answer, answer.Answer, "answer answer should be equal to the new question answer")
		assert.NotEmpty(t, answer.Hint, "answer hint should not be empty")
	})
}

func TestHub(t *testing.T) {
	t.Run("should create a new hub", func(t *testing.T) {
		hub := tests.RunningHub(t)
		assert.NotEmpty(t, hub.Controller, "hub id should not be empty")
		assert.Empty(t, hub.Clients, "hub clients should be empty")
	})

	t.Run("should create a new hub with 4 clients", func(t *testing.T) {
		hub := tests.RunningHubWithClients(t, 4)
		assert.NotEmpty(t, hub.Controller, "hub id should not be empty")
		assert.Equal(t, 4, len(hub.Clients), "hub clients should be equal to 4")
	})

	t.Run("should create a new client", func(t *testing.T) {
		hub := tests.RunningHubWithClients(t, 4)
		client := tests.NewClient(t, hub)
		assert.Empty(t, client.PlayerId, "client id should not be empty")
		assert.Equal(t, hub, client.Hub, "client hub should be equal to the hub")
	})

	t.Run("should find client x from hub", func(t *testing.T) {
		hub := tests.RunningHubWithClients(t, 4)
		client := tests.ClientX(hub)
		assert.Empty(t, client.PlayerId, "client id should not be empty")
		assert.Equal(t, hub, client.Hub, "client hub should be equal to the hub")
	})

	t.Run("should find client y from hub", func(t *testing.T) {
		hub := tests.RunningHub(t)
		client := tests.ClientX(hub)
		assert.Nil(t, client, "client should not be nil")
	})
}

func TestMessage(t *testing.T) {
	t.Run("should create a new message", func(t *testing.T) {
		message := tests.TestMessage(actions.ACTIVE, "this is a test message")
		assert.Equal(t, "ACTIVE", message.Action, "message action should be equal to ACTIVE")
		assert.Equal(t, "this is a test message", message.Content, "message content should be equal to 'this is a test message'")
	})

	t.Run("should create a new begin message", func(t *testing.T) {
		player := tests.TestPlayer()
		message := tests.TestBeginMessage(player)
		assert.Equal(t, "BEGIN", message.Action, "message action should be equal to BEGIN")
		assert.NotEmpty(t, message.Content, "message content should not be empty")
	})
}

func TestSubscriber(t *testing.T) {
	t.Run("should create a new subscriber", func(t *testing.T) {
		subscriber := tests.TestSubscriber()
		assert.NotEmpty(t, subscriber.Tag, "subscriber tag should not be empty")
		assert.NotEmpty(t, subscriber.PlayerId, "subscriber player id should not be empty")
		assert.Equal(t, "PLAYER", subscriber.Role, "subscriber role should be equal to PLAYER")
		assert.True(t, subscriber.Active, "subscriber should be active after initialization")
	})
}

func TestSnapshot(t *testing.T) {
	t.Run("should create a new begin message", func(t *testing.T) {
		quiz := tests.TestGame()
		snapshot := tests.TestSnapshot(quiz.Id)
		assert.Equal(t, quiz.Id, snapshot.QuizId, "snapshot quiz id must match")
		assert.Empty(t, snapshot.Roster, "snapshot roster must be empty")
		assert.Equal(t, "CREATE", snapshot.EventType, "snapshot event type must be CREATE")
	})
}
