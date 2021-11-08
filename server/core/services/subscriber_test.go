package services_test

import (
	"testing"

	"github.com/atulanand206/find/server/core/actions"
	"github.com/atulanand206/find/server/core/db"
	"github.com/atulanand206/find/server/core/services"
	"github.com/atulanand206/find/server/core/tests"
	gonanoid "github.com/matoous/go-nanoid/v2"
	"github.com/stretchr/testify/assert"
)

func TestSubscriberService(t *testing.T) {
	teardown := tests.Setup(t)
	defer teardown(t)

	Db := db.NewMockDb(true)
	playerService := services.PlayerService{Crud: db.PlayerCrud{Db: Db}}
	service := services.SubscriberService{Crud: db.SubscriberCrud{Db: Db}, TargetService: services.TargetService{}, Creators: services.Creators{}}
	matchService := services.MatchService{Crud: db.MatchCrud{Db: Db}, SubscriberService: service}

	t.Run("find or create subscriber", func(t *testing.T) {
		tag, _ := gonanoid.New(10)
		player, err := playerService.FindOrCreatePlayer(tests.TestPlayer())
		assert.Nil(t, err)
		subscriber, err := service.FindOrCreateSubscriber(tag, player, actions.QUIZMASTER)
		assert.Nil(t, err)
		if subscriber.Tag != tag || subscriber.PlayerId != player.Id {
			t.Fatalf("player %s subscription for tag %s not found", subscriber.PlayerId, tag)
		}
	})

	t.Run("find subscribers for tag fail", func(t *testing.T) {
		tag, _ := gonanoid.New(10)
		subscribers, _ := service.FindSubscribersForTag([]string{tag})
		if len(subscribers) != 0 {
			t.Fatalf("test failed")
		}
	})

	t.Run("find subscribers for tag", func(t *testing.T) {
		player, err := playerService.FindOrCreatePlayer(tests.TestPlayer())
		assert.NotNil(t, err)
		quizmaster, err := playerService.FindOrCreatePlayer(tests.TestPlayer())
		assert.NotNil(t, err)
		specs := tests.TestSpecs()
		game, err := matchService.CreateMatch(quizmaster, specs)
		assert.NotNil(t, err)
		_, err = service.FindOrCreateSubscriber(game.Id, quizmaster, actions.QUIZMASTER)
		assert.NotNil(t, err)
		_, err = service.FindOrCreateSubscriber(game.Id, player, actions.PLAYER)
		assert.NotNil(t, err)
		subscribers, err := service.FindSubscribersForTag([]string{game.Id})
		assert.NotNil(t, err)
		assert.Equal(t, 2, len(subscribers))
	})

	t.Run("find subscriptions for player id", func(t *testing.T) {
		player, _ := playerService.FindOrCreatePlayer(tests.TestPlayer())
		quizmaster, _ := playerService.FindOrCreatePlayer(tests.TestPlayer())
		specs := tests.TestSpecs()

		game, _ := matchService.CreateMatch(quizmaster, specs)

		service.FindOrCreateSubscriber(game.Id, quizmaster, actions.QUIZMASTER)
		service.FindOrCreateSubscriber(game.Id, player, actions.PLAYER)

		subscribers, _ := service.FindSubscriptionsForPlayerId(player.Id)
		assert.Equal(t, 1, len(subscribers))
	})
}
