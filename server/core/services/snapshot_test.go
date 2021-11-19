package services_test

import (
	"testing"

	"github.com/atulanand206/find/server/core/db"
	"github.com/atulanand206/find/server/core/models"
	"github.com/atulanand206/find/server/core/services"
	"github.com/atulanand206/find/server/core/tests"
	"github.com/atulanand206/find/server/core/utils"
	"github.com/stretchr/testify/assert"
)

func TestSnapshotService(t *testing.T) {
	teardown := tests.Setup(t)
	defer teardown(t)

	service := services.SnapshotService{db.SnapshotCrud{Db: db.NewMockDb(true)}}

	t.Run("add snapshot", func(t *testing.T) {
		quiz := tests.TestGame()
		teams := tests.TestTeams(quiz.Id, 2)
		teamRoster := utils.TableRoster(teams, []models.Subscriber{}, []models.Player{})
		snapshot := models.InitialSnapshot(quiz.Id, teamRoster)

		err := service.CreateSnapshot(snapshot)
		assert.Nil(t, err, "error must be nil")
	})

	t.Run("find snapshot for match without creation", func(t *testing.T) {
		snapshots, err := service.FindSnapshotsForMatch(tests.TestRandomString())
		assert.Nil(t, err, "error must be nil")
		assert.Equal(t, 0, len(snapshots), "snapshots must be empty")
	})

	t.Run("find snapshot for match", func(t *testing.T) {
		quiz := tests.TestGame()
		teams := tests.TestTeams(quiz.Id, 2)
		teamRoster := utils.TableRoster(teams, []models.Subscriber{}, []models.Player{})
		snapshot := models.InitialSnapshot(quiz.Id, teamRoster)

		err := service.CreateSnapshot(snapshot)
		assert.Nil(t, err, "error must be nil")

		snapshots, err := service.FindSnapshotsForMatch(quiz.Id)
		assert.Nil(t, err, "error must be nil")
		assert.Equal(t, 1, len(snapshots), "snapshots count must be one")
		assert.Equal(t, snapshot, snapshots[0], "snapshot must match")
	})
}
