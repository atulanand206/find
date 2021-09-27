package core

import (
	"fmt"

	"github.com/google/uuid"
)

func InitNewIndex(tag string) (index Index) {
	index.Tag = tag
	index.Id = id(index)
	return
}

func InitNewQuestion(index Index, newQuestion NewQuestion) (question Question) {
	question.Statements = newQuestion.Statements
	question.Tag = index.Id
	question.Id = id(question)
	return
}

func InitNewAnswer(question Question, newQuestion NewQuestion) (answer Answer) {
	answer.Answer = newQuestion.Answer
	answer.QuestionId = question.Id
	answer.Id = id(answer)
	return
}

func InitNewMatch(quizmaster Player) (match Game) {
	match.Teams = make([]Team, 0)
	for idx := 0; idx < TeamsInAMatch; idx++ {
		match.Teams = append(match.Teams, InitNewTeam(NewTeamName(idx)))
	}
	match.QuizMaster = quizmaster
	match.Id = id(match)
	return
}

func NewTeamName(idx int) (name string) {
	return fmt.Sprintf("Team %d", idx)
}

func InitNewTeam(name string) (team Team) {
	team.Players = make([]Player, 0)
	team.Name = name
	team.Id = id(team)
	return
}

func InitNewPlayer(playerRequest Player) (player Player) {
	player.Name = playerRequest.Name
	player.Email = playerRequest.Email
	player.Id = id(player)
	return
}

func id(v interface{}) string {
	return uuid.NewString()
}

func InitFindAnswerResponse(questionId string, answer Answer) (response FindAnswerResponse) {
	response.QuestionId = questionId
	response.Answer = answer.Answer
	return
}

func InitAddQuestionResponse(question Question, answer Answer) (response AddQuestionResponse) {
	response.QuestionId = question.Id
	response.AnswerId = answer.Id
	return
}

func InitStartGameResponse(match Game, questions []Question) (response StartGameResponse) {
	response.Match = match
	response.Prompt = questions
	return
}

func InitWebSocketMessageFailure() (response WebsocketMessage) {
	response = InitWebSocketMessage(Failure, Err_SocketRequestFailed)
	return
}

func InitWebSocketMessage(action Action, content string) (response WebsocketMessage) {
	response.Action = action
	response.Content = content
	return
}
