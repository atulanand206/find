package core_test

import (
	"testing"

	"github.com/atulanand206/find/server/core"
	gonanoid "github.com/matoous/go-nanoid/v2"
)

func TestFindOrCreateSubscriber(t *testing.T) {
	teardown := Setup(t)
	defer teardown(t)

	service := core.SubscriberService{}

	tag, _ := gonanoid.New(10)
	player := testPlayer()
	subscriber, err := service.FindOrCreateSubscriber(tag, player, core.QUIZMASTER)
	if err != nil {
		t.Fatalf("Error creating subscriber: %v", err)
	}
	if subscriber.Tag != tag || subscriber.PlayerId != player.Id {
		t.Fatalf("player %s subscription for tag %s not found", subscriber.PlayerId, tag)
	}
}

func TestFindSubscribersForTagFail(t *testing.T) {
	teardown := Setup(t)
	defer teardown(t)

	service := core.SubscriberService{}

	tag, _ := gonanoid.New(10)
	subscribers, _ := service.FindSubscribersForTag([]string{tag})
	if len(subscribers) != 0 {
		t.Fatalf("test failed")
	}
}

func TestFindSubscribersForTag(t *testing.T) {
	teardown := Setup(t)
	defer teardown(t)

	playerService := core.PlayerService{}
	player, _ := playerService.FindOrCreatePlayer(testPlayer())
	quizmaster, _ := playerService.FindOrCreatePlayer(testPlayer())
	specs := testSpecs()

	matchService := core.MatchService{}
	game, _ := matchService.CreateMatch(quizmaster, specs)

	subscriberService := core.SubscriberService{}
	subscriberService.FindOrCreateSubscriber(game.Id, quizmaster, core.QUIZMASTER)
	subscriberService.FindOrCreateSubscriber(game.Id, player, core.PLAYER)

	subscribers, _ := subscriberService.FindSubscribersForTag([]string{game.Id})
	if len(subscribers) != 2 {
		t.Fatalf("test failed")
	}
}

func TestFindSubscriptionsForPlayerId(t *testing.T) {
	teardown := Setup(t)
	defer teardown(t)

	playerService := core.PlayerService{}
	player, _ := playerService.FindOrCreatePlayer(testPlayer())
	quizmaster, _ := playerService.FindOrCreatePlayer(testPlayer())
	specs := testSpecs()

	matchService := core.MatchService{}
	game, _ := matchService.CreateMatch(quizmaster, specs)

	subscriberService := core.SubscriberService{}
	subscriberService.FindOrCreateSubscriber(game.Id, quizmaster, core.QUIZMASTER)
	subscriberService.FindOrCreateSubscriber(game.Id, player, core.PLAYER)

	subscribers, _ := subscriberService.FindSubscriptionsForPlayerId(player.Id)
	if len(subscribers) != 1 {
		t.Fatalf("test failed")
	}
}
