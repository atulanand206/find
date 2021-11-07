package comms

import (
	"encoding/json"
	"net/http"

	"github.com/atulanand206/find/server/core/errors"
	"github.com/atulanand206/find/server/core/services"
)

type MatchHandler struct {
	matchService services.MatchService
}

func (handler MatchHandler) HandlerActiveQuizzes(w http.ResponseWriter, r *http.Request) {
	var request CreatePermissionRequest
	if err := DecodeJsonBody(r, &request); err != nil {
		http.Error(w, errors.Err_RequestNotDecoded, http.StatusInternalServerError)
		return
	}

	matches, err := handler.matchService.FindActiveMatchesForPlayer(request.PlayerId)
	if err != nil || len(matches) == 0 {
		er := Controller.Creators.ErrorCreator.ActiveMatchesNotFound()
		http.Error(w, er.Msg, er.Code)
		return
	}

	json.NewEncoder(w).Encode(matches)
}
