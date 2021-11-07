package services

import "github.com/atulanand206/find/server/core/db"

type PermissionService struct {
	Crud db.PermissionCrud
}

func (servie PermissionService) HasPermissions(playerId string) bool {
	return servie.Crud.HasPermission(playerId)
}
