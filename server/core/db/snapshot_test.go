package db_test

import (
	"testing"

	"github.com/atulanand206/find/server/core/db"
	"github.com/atulanand206/find/server/core/tests"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
)

func TestSnapshot(t *testing.T) {
	teardown := tests.Setup(t)
	defer teardown(t)

	crud := db.SnapshotCrud{Db: db.NewDb()}

	t.Run("should create a new snapshot", func(t *testing.T) {
		quiz := tests.TestGame()
		snapshot := tests.TestSnapshot(quiz.Id)
		err := crud.CreateSnapshot(snapshot)
		assert.Nil(t, err, "error must be nil")
	})

	t.Run("should create and find match snapshots", func(t *testing.T) {
		quiz := tests.TestGame()
		snapshot := tests.TestSnapshot(quiz.Id)
		err := crud.CreateSnapshot(snapshot)
		assert.Nil(t, err, "error must be nil")
		foundSnapshots, err := crud.FindSnapshots(bson.M{"quiz_id": quiz.Id})
		assert.Nil(t, err, "error must be nil")
		assert.Equal(t, 1, len(foundSnapshots), "must find one snapshot")
		assert.Equal(t, snapshot, foundSnapshots[0], "subscriber must be equal")
	})

	t.Run("should create and find latest snapshot", func(t *testing.T) {
		quiz := tests.TestGame()
		snapshot := tests.TestSnapshot(quiz.Id)
		err := crud.CreateSnapshot(snapshot)
		assert.Nil(t, err, "error must be nil")
		foundSnapshot, err := crud.FindLatestSnapshot(quiz.Id)
		assert.Nil(t, err, "error must be nil")
		assert.Equal(t, snapshot, foundSnapshot, "snapshot must be equal")
	})

	t.Run("fail to find latests snapshot", func(t *testing.T) {
		_, err := crud.FindLatestSnapshot(tests.TestRandomString())
		assert.NotNil(t, err, "error must not be nil")
	})
}
