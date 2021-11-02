package core

import (
	"encoding/json"
	"net/http"
)

type MatchHandler struct {
	matchService MatchService
}

func (handler MatchHandler) HandlerActiveQuizzes(w http.ResponseWriter, r *http.Request) {
	var request CreatePermissionRequest
	if err := DecodeJsonBody(r, &request); err != nil {
		http.Error(w, Err_RequestNotDecoded, http.StatusInternalServerError)
		return
	}

	matches, err := handler.matchService.FindActiveMatchesForPlayer(request.PlayerId)
	if err != nil || len(matches) == 0 {
		er := ErrorCreator.ActiveMatchesNotFound()
		http.Error(w, er.msg, er.code)
		return
	}

	json.NewEncoder(w).Encode(matches)
}
