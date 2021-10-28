package core

import (
	"fmt"
)

type Validator struct{}

func (validator Validator) ValidateRequest(request Request) (err error) {
	switch request.Action {
	case PLAYER.String():
		return validator.ValidateBeginRequest(request)
	case SPECS.String():
		return validator.ValidateCreateQuizRequest(request)
	case REFRESH.String():
		return validator.ValidateRefeshQuizRequest(request)
	}
	return
}
func (validator Validator) ValidateBeginRequest(request Request) (err error) {
	return
}

func (validator Validator) ValidateCreateQuizRequest(request Request) (err error) {
	fmt.Println(request)
	// && request.Person != nil && request.Specs != nil
	return
}

func (validator Validator) ValidateRefeshQuizRequest(request Request) (err error) {
	return
}
