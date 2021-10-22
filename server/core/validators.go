package core

import (
	"errors"
	"fmt"
)

type Validator struct{}

func (validator Validator) ValidateBeginRequest(request Request) (err error) {
	return
}

func (validator Validator) ValidateCreateQuizRequest(request Request) (err error) {
	fmt.Println(request)
	// && request.Person != nil && request.Specs != nil
	return
}

func (validator Validator) ValidateRefeshQuizRequest(request Request) (err error) {
	if len(request.QuizId) != 0 {
		err = errors.New("request invalid.")
	}
	return
}
