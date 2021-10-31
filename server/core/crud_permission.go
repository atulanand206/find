package core

import (
	"github.com/atulanand206/go-mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const PermissionCollection = "permissions"

type PermissionCrud struct {
	db DB
}

func (crud PermissionCrud) CreatePermission(playerId string) (err error) {
	if crud.HasPermission(playerId) {
		return
	}
	return crud.db.Create(Permission{PlayerId: playerId}, PermissionCollection)
}

func (crud PermissionCrud) HasPermission(playerId string) bool {
	item := mongo.FindOne(Database, PermissionCollection, bson.M{"player_id": playerId}, &options.FindOneOptions{})
	return item != nil
}

func (crud PermissionCrud) FindPermissions() (permissions []Permission, err error) {
	cursor, err := mongo.Find(Database, PermissionCollection, bson.M{}, &options.FindOptions{})
	if err != nil {
		return
	}
	permissions, err = DecodePermissions(cursor)
	return
}
