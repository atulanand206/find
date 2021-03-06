package db

import (
	"github.com/atulanand206/find/server/core/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type PermissionCrud struct {
	Db DB
}

func (crud PermissionCrud) CreatePermission(playerId string) (err error) {
	if crud.HasPermission(playerId) {
		return
	}
	return crud.Db.Create(models.Permission{PlayerId: playerId}, PermissionCollection)
}

func (crud PermissionCrud) HasPermission(playerId string) bool {
	cursor, err := crud.Db.Find(PermissionCollection, bson.M{"player_id": playerId}, &options.FindOptions{})
	if err != nil {
		return false
	}
	permissions, err := DecodePermissions(cursor)
	if err != nil || len(permissions) == 0 {
		return false
	}
	return true
}

func (crud PermissionCrud) FindPermissions() (permissions []models.Permission, err error) {
	cursor, err := crud.Db.Find(PermissionCollection, bson.M{}, &options.FindOptions{})
	if err != nil {
		return
	}
	permissions, err = DecodePermissions(cursor)
	return
}
