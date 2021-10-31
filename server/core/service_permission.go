package core

type PermissionService struct {
	crud PermissionCrud
}

func (servie PermissionService) HasPermissions(playerId string) bool {
	return servie.crud.HasPermission(playerId)
}
