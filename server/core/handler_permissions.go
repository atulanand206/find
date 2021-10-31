package core

import (
	"encoding/json"
	"net/http"
)

type CreatePermissionRequest struct {
	PlayerId string `json:"player_id"`
}

type PermissionHandler struct {
	crud PermissionCrud
}

func (handler PermissionHandler) HandlerCreatePermission(w http.ResponseWriter, r *http.Request) {
	var request CreatePermissionRequest
	if err := DecodeJsonBody(r, &request); err != nil {
		http.Error(w, Err_RequestNotDecoded, http.StatusInternalServerError)
		return
	}

	if err := handler.crud.CreatePermission(request.PlayerId); err != nil {
		er := ErrorCreator.PermissionNotCreated(request.PlayerId)
		http.Error(w, er.msg, er.code)
		return
	}
}

func DecodeJsonBody(r *http.Request, v interface{}) error {
	decoder := json.NewDecoder(r.Body)
	return decoder.Decode(v)
}

func (handler PermissionHandler) HandlerFindPermissions(w http.ResponseWriter, r *http.Request) {
	permissions, err := handler.crud.FindPermissions()
	if err != nil || len(permissions) == 0 {
		er := ErrorCreator.PermissionsNotFound()
		http.Error(w, er.msg, er.code)
		return
	}

	json.NewEncoder(w).Encode(permissions)
}
