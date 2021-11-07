package services

import (
	"fmt"

	"github.com/atulanand206/find/server/core/actions"
	"github.com/atulanand206/find/server/core/errors"
	"github.com/atulanand206/find/server/core/models"
)

type Validator struct{}

func (validator Validator) ValidateRequest(request models.Request) (err error) {
	switch request.Action {
	case actions.BEGIN.String():
		return validator.ValidateBeginRequest(request)
	case actions.SPECS.String():
		return validator.ValidateCreateQuizRequest(request)
	case actions.REFRESH.String():
		return validator.ValidateRefeshQuizRequest(request)
	}
	return
}
func (validator Validator) ValidateBeginRequest(request models.Request) (err error) {
	if request.Person.Email == "" || request.Person.Name == "" {
		err = fmt.Errorf(errors.Err_RequestNotValid)
	}
	return
}

func (validator Validator) ValidateCreateQuizRequest(request models.Request) (err error) {
	if request.Specs.Name == "" ||
		request.Specs.Players < 0 ||
		request.Specs.Teams < 0 || request.Specs.Teams > 4 ||
		request.Specs.Rounds < 0 ||
		request.Specs.Questions < 0 || request.Specs.Questions > 20 ||
		request.Specs.Points < 0 {
		err = fmt.Errorf(errors.Err_RequestNotValid)
	}
	return
}

func (validator Validator) ValidateRefeshQuizRequest(request models.Request) (err error) {
	return
}
