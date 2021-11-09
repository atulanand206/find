package db_test

import (
	"testing"

	"github.com/atulanand206/find/server/core/db"
	"github.com/atulanand206/find/server/core/tests"
	"github.com/stretchr/testify/assert"
)

func TestPermissionCrud(t *testing.T) {
	teardown := tests.Setup(t)
	defer teardown(t)

	crud := db.PermissionCrud{Db: db.NewDb()}

	t.Run("create permission", func(t *testing.T) {
		playerId := tests.TestId()
		err := crud.CreatePermission(playerId)
		assert.Nil(t, err, "error must be nil")
	})

	t.Run("has permission", func(t *testing.T) {
		playerId := tests.TestId()
		err := crud.CreatePermission(playerId)
		assert.Nil(t, err, "error must be nil")
		present := crud.HasPermission(playerId)
		assert.True(t, present, "player must have permission")
	})

	t.Run("find permissions", func(t *testing.T) {
		crud.Db.DropCollections()
		playerId1 := tests.TestId()
		err := crud.CreatePermission(playerId1)
		assert.Nil(t, err, "error must be nil")

		playerId2 := tests.TestId()
		err = crud.CreatePermission(playerId2)
		assert.Nil(t, err, "error must be nil")

		permissions, err := crud.FindPermissions()
		assert.Nil(t, err, "error must be nil")
		assert.Equal(t, 2, len(permissions))
	})
}
