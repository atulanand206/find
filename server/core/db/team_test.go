package db_test

import (
	"testing"

	"github.com/atulanand206/find/server/core/db"
	"github.com/atulanand206/find/server/core/tests"
	"github.com/stretchr/testify/assert"
)

func TestTeamCrud(t *testing.T) {
	teardown := tests.Setup(t)
	defer teardown(t)

	crud := db.TeamCrud{Db: db.NewDb()}

	t.Run("create teams", func(t *testing.T) {
		quizId := tests.TestRandomString()
		teams := tests.TestTeams(quizId, 2)
		err := crud.CreateTeams(teams)
		assert.Nil(t, err, "error must be nil")
	})

	t.Run("create and find teams", func(t *testing.T) {
		quizId := tests.TestRandomString()
		teams := tests.TestTeams(quizId, 2)
		err := crud.CreateTeams(teams)
		assert.Nil(t, err, "error must be nil")
		foundTeams, err := crud.FindTeams(quizId)
		assert.Nil(t, err, "error must be nil")
		assert.Equal(t, len(teams), len(foundTeams), "number of teams must be equal")
	})

	t.Run("create and update teams", func(t *testing.T) {
		quizId := tests.TestRandomString()
		teams := tests.TestTeams(quizId, 2)
		err := crud.CreateTeams(teams)
		assert.Nil(t, err, "error must be nil")
		team := teams[0]
		team.Name = "new name"
		err = crud.UpdateTeam(team)
		assert.Nil(t, err, "error must be nil")
		foundTeams, err := crud.FindTeams(quizId)
		assert.Nil(t, err, "error must be nil")
		assert.Equal(t, len(teams), len(foundTeams), "number of teams must be equal")
		for _, foundTeam := range foundTeams {
			if foundTeam.Id == team.Id {
				assert.Equal(t, team.Name, foundTeam.Name, "team name must be equal")
			}
		}
	})
}
