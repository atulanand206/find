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

func InitStartGameResponse(quizId string, teams []Team, question Question, snapshot Snapshot) (response StartGameResponse) {
	response.QuizId = quizId
	response.Teams = teams
	response.Snapshot = snapshot
	return
}

func InitScoreResponse(request ScoreRequest, snapshots []Snapshot) (response ScoreResponse) {
	response.QuizId = request.QuizId
	response.Snapshots = snapshots
	return
}

func InitSnapshotDtoF(quizId string,
	questionId string,
	teamsTurn string, eventType string,
	score int, questionNo int, roundNo int,
	content []string) (response Snapshot) {
	response.QuizId = quizId
	response.QuestionId = questionId
	response.TeamSTurn = teamsTurn
	response.EventType = eventType
	response.Content = content
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
