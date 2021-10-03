package core

import (
	"fmt"
	"time"

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

func InitNewMatch(quizmaster Player, specs Specs) (match Game) {
	match.Teams = make([]TeamMini, 0)
	for idx := 0; idx < specs.Teams; idx++ {
		match.Teams = append(match.Teams, InitNewTeam(NewTeamName(idx), specs.Players))
	}
	match.Tags = make([]string, 0)
	match.QuizMaster = quizmaster
	match.Specs = specs
	match.Specs.Points = 16
	match.Specs.Rounds = 2
	match.Id = id(match)
	return
}

func NewTeamName(idx int) (name string) {
	return fmt.Sprintf("Team %d", idx)
}

func InitNewTeam(name string, players int) (team TeamMini) {
	team.Name = name
	team.Id = id(team)
	return
}

func InitNewTeamM(teammini TeamMini) (team Team) {
	team.Players = make([]Player, 0)
	team.Name = teammini.Name
	team.Id = teammini.Id
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

func InitEnterGameResponse(match Game, teams []Team) (response EnterGameResponse) {
	response.Quiz = match
	response.Teams = teams
	return
}

func InitStartGameResponse(quizId string, teams []Team, questions []Question) (response StartGameResponse) {
	response.QuizId = quizId
	response.Teams = teams
	response.Question = questions[0]
	response.TeamSTurn = teams[0].Id
	response.QuestionNo = 1
	response.RoundNo = 1
	return
}

func InitHintRevealResponse(request GameSnapRequest, answer Answer, questionNo int, roundNo int) (response HintRevealResponse) {
	response.QuizId = request.QuizId
	response.QuestionId = request.QuestionId
	response.TeamSTurn = request.TeamSTurn
	response.Hint = answer.Hint
	response.QuestionNo = questionNo
	response.RoundNo = roundNo
	return
}

func InitAnswerRevealResponse(request GameSnapRequest, answer Answer, questionNo int, roundNo int) (response AnswerRevealResponse) {
	response.QuizId = request.QuizId
	response.QuestionId = request.QuestionId
	response.TeamSTurn = request.TeamSTurn
	response.Answer = answer.Answer
	response.QuestionNo = questionNo
	response.RoundNo = roundNo
	return
}

func InitNextQuestionResponse(request GameSnapRequest, question Question, teamsTurn string, questionNo int, roundNo int) (response GameNextResponse) {
	response.QuizId = request.QuizId
	response.LastQuestionId = request.QuestionId
	response.TeamSTurn = teamsTurn
	response.Question = question
	response.QuestionNo = questionNo
	response.RoundNo = roundNo
	return
}

func InitPassQuestionResponse(request GameSnapRequest, teamsTurn string, questionNo int, roundNo int) (response GamePassResponse) {
	response.QuizId = request.QuizId
	response.TeamSTurn = teamsTurn
	response.QuestionId = request.QuestionId
	response.QuestionNo = questionNo
	response.RoundNo = roundNo
	return
}

func InitScoreResponse(request ScoreRequest, snapshots []Snapshot) (response ScoreResponse) {
	response.QuizId = request.QuizId
	response.Snapshots = snapshots
	return
}

func InitSnapshotDto(request GameSnapRequest, eventType string, score int, questionNo int, roundNo int) (response Snapshot) {
	response.QuizId = request.QuizId
	response.QuestionId = request.QuestionId
	response.TeamSTurn = request.TeamSTurn
	response.EventType = eventType
	response.Score = score
	response.QuestionNo = questionNo
	response.RoundNo = roundNo
	response.Timestamp = time.Now().String()
	return
}

func InitSnapshotDtoF(quizId string, questionId string, teamsTurn string, eventType string, score int, questionNo int, roundNo int) (response Snapshot) {
	response.QuizId = quizId
	response.QuestionId = questionId
	response.TeamSTurn = teamsTurn
	response.EventType = eventType
	response.Score = score
	response.QuestionNo = questionNo
	response.RoundNo = roundNo
	response.Timestamp = time.Now().String()
	return
}

func InitWebSocketMessageFailure() (response WebsocketMessage) {
	response = InitWebSocketMessage(Failure, Err_SocketRequestFailed)
	return
}

func InitWebSocketMessage(action Action, content string) (response WebsocketMessage) {
	response.Action = action.String()
	response.Content = content
	return
}
