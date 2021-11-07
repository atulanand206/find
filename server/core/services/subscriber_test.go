package services_test

import (
	"testing"

	"github.com/atulanand206/find/server/core/actions"
	"github.com/atulanand206/find/server/core/services"
	"github.com/atulanand206/find/server/core/tests"
	gonanoid "github.com/matoous/go-nanoid/v2"
)

func TestFindOrCreateSubscriber(t *testing.T) {
	teardown := tests.Setup(t)
	defer teardown(t)

	service := services.SubscriberService{}

	tag, _ := gonanoid.New(10)
	player := tests.TestPlayer()
	subscriber, err := service.FindOrCreateSubscriber(tag, player, actions.QUIZMASTER)
	if err != nil {
		t.Fatalf("Error creating subscriber: %v", err)
	}
	if subscriber.Tag != tag || subscriber.PlayerId != player.Id {
		t.Fatalf("player %s subscription for tag %s not found", subscriber.PlayerId, tag)
	}
}

func TestFindSubscribersForTagFail(t *testing.T) {
	teardown := tests.Setup(t)
	defer teardown(t)

	service := services.SubscriberService{}

	tag, _ := gonanoid.New(10)
	subscribers, _ := service.FindSubscribersForTag([]string{tag})
	if len(subscribers) != 0 {
		t.Fatalf("test failed")
	}
}

func TestFindSubscribersForTag(t *testing.T) {
	teardown := tests.Setup(t)
	defer teardown(t)

	playerService := services.PlayerService{}
	player, _ := playerService.FindOrCreatePlayer(tests.TestPlayer())
	quizmaster, _ := playerService.FindOrCreatePlayer(tests.TestPlayer())
	specs := tests.TestSpecs()

	matchService := services.MatchService{}
	game, _ := matchService.CreateMatch(quizmaster, specs)

	subscriberService := services.SubscriberService{}
	subscriberService.FindOrCreateSubscriber(game.Id, quizmaster, actions.QUIZMASTER)
	subscriberService.FindOrCreateSubscriber(game.Id, player, actions.PLAYER)

	subscribers, _ := subscriberService.FindSubscribersForTag([]string{game.Id})
	if len(subscribers) != 2 {
		t.Fatalf("test failed")
	}
}

func TestFindSubscriptionsForPlayerId(t *testing.T) {
	teardown := tests.Setup(t)
	defer teardown(t)

	playerService := services.PlayerService{}
	player, _ := playerService.FindOrCreatePlayer(tests.TestPlayer())
	quizmaster, _ := playerService.FindOrCreatePlayer(tests.TestPlayer())
	specs := tests.TestSpecs()

	matchService := services.MatchService{}
	game, _ := matchService.CreateMatch(quizmaster, specs)

	subscriberService := services.SubscriberService{}
	subscriberService.FindOrCreateSubscriber(game.Id, quizmaster, actions.QUIZMASTER)
	subscriberService.FindOrCreateSubscriber(game.Id, player, actions.PLAYER)

	subscribers, _ := subscriberService.FindSubscriptionsForPlayerId(player.Id)
	if len(subscribers) != 1 {
		t.Fatalf("test failed")
	}
}
