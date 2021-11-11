package db_test

import (
	"testing"

	"github.com/atulanand206/find/server/core/db"
	"github.com/atulanand206/find/server/core/tests"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
)

func TestSubscriber(t *testing.T) {
	teardown := tests.Setup(t)
	defer teardown(t)

	crud := db.SubscriberCrud{Db: db.NewDb()}

	t.Run("should create a new subscriber", func(t *testing.T) {
		subscriber := tests.TestSubscriber()
		err := crud.CreateSubscriber(subscriber)
		assert.Nil(t, err, "error must be nil")
	})

	t.Run("should create and find a subscriber", func(t *testing.T) {
		subscriber := tests.TestSubscriber()
		err := crud.CreateSubscriber(subscriber)
		assert.Nil(t, err, "error must be nil")
		foundSubscriber, err := crud.FindSubscriber(subscriber.Tag, subscriber.PlayerId)
		assert.Nil(t, err, "error must be nil")
		assert.Equal(t, subscriber, foundSubscriber, "subscriber must be equal")
	})

	t.Run("should create and find subscriptions", func(t *testing.T) {
		subscriber := tests.TestSubscriber()
		err := crud.CreateSubscriber(subscriber)
		assert.Nil(t, err, "error must be nil")
		foundSubscribers, err := crud.FindSubscribers(bson.M{"player_id": subscriber.PlayerId})
		assert.Nil(t, err, "error must be nil")
		assert.Equal(t, 1, len(foundSubscribers), "subscriptions must be one")
		assert.Equal(t, subscriber, foundSubscribers[0], "subscriber must be equal")
	})

	t.Run("fail to find subscriber", func(t *testing.T) {
		subscriber := tests.TestSubscriber()
		_, err := crud.FindSubscriber(subscriber.Tag, subscriber.PlayerId)
		assert.NotNil(t, err, "error must not be nil")
	})

	t.Run("should delete subscription", func(t *testing.T) {
		subscriber := tests.TestSubscriber()
		err := crud.CreateSubscriber(subscriber)
		assert.Nil(t, err, "error must be nil")
		foundSubscriber, err := crud.FindSubscriber(subscriber.Tag, subscriber.PlayerId)
		assert.Nil(t, err, "error must be nil")
		assert.Equal(t, subscriber, foundSubscriber, "subscriber must be equal")
		err = crud.DeleteSubscriber(subscriber.PlayerId)
		assert.Nil(t, err, "error must be nil")
		_, err = crud.FindSubscriber(subscriber.Tag, subscriber.PlayerId)
		assert.NotNil(t, err, "error must not be nil")
	})
}
