package services

import (
	"fmt"

	"github.com/atulanand206/find/server/core/actions"
	"github.com/atulanand206/find/server/core/models"
)

type Validator struct{}

func (validator Validator) ValidateRequest(request models.Request) (err error) {
	switch request.Action {
	case actions.PLAYER.String():
		return validator.ValidateBeginRequest(request)
	case actions.SPECS.String():
		return validator.ValidateCreateQuizRequest(request)
	case actions.REFRESH.String():
		return validator.ValidateRefeshQuizRequest(request)
	}
	return
}
func (validator Validator) ValidateBeginRequest(request models.Request) (err error) {
	return
}

func (validator Validator) ValidateCreateQuizRequest(request models.Request) (err error) {
	fmt.Println(request)
	// && request.Person != nil && request.Specs != nil
	return
}

func (validator Validator) ValidateRefeshQuizRequest(request models.Request) (err error) {
	return
}
