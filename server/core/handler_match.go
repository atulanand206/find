package core

import (
	"encoding/json"
	"net/http"
)

type MatchHandler struct {
	matchService MatchService
}

func (handler MatchHandler) HandlerActiveQuizzes(w http.ResponseWriter, r *http.Request) {
	matches, err := handler.matchService.FindActiveMatches()
	if err != nil || len(matches) == 0 {
		er := ErrorCreator.ActiveMatchesNotFound()
		http.Error(w, er.msg, er.code)
		return
	}

	json.NewEncoder(w).Encode(matches)
}
