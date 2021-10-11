package core

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	gonanoid "github.com/matoous/go-nanoid/v2"
)

type Creator struct{}

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

func InitNewTeams(match Game) (teams []Team) {
	teams = make([]Team, 0)
	for idx := 0; idx < match.Specs.Teams; idx++ {
		teams = append(teams, InitNewTeam(NewTeamName(idx), match.Id))
	}
	return
}

func InitNewMatch(quizmaster Player, specs Specs) (match Game) {
	match.Tags = make([]string, 0)
	match.QuizMaster = quizmaster
	match.Specs = specs
	match.Specs.Points = 16
	match.Specs.Rounds = 2
	id, _ := gonanoid.New(8)
	match.Id = id
	return
}

func NewTeamName(idx int) (name string) {
	return fmt.Sprintf("Team %d", idx)
}

func InitNewTeam(name string, quizId string) (team Team) {
	team.Name = name
	team.QuizId = quizId
	team.Id = id(team)
	return
}

func InitNewPlayer(playerRequest Player) (player Player) {
	player.Name = playerRequest.Name
	player.Email = playerRequest.Email
	player.Id = playerRequest.Id
	return
}

func (creator Creator) InitSubscriber(game Game, player Player, role string) (subscriber Subscriber) {
	subscriber.Tag = game.Id
	subscriber.PlayerId = player.Id
	subscriber.Role = role
	subscriber.Active = true
	return
}

func InitTeamPlayer(teamId string, player Player) (teamPlayer TeamPlayerRequest) {
	teamPlayer.TeamId = teamId
	teamPlayer.PlayerId = player.Id
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

func InitEnterGameResponse(match Game, teams []Team, teamPlayers []TeamPlayer, players []Player, playerTeamId string, snapshot Snapshot) (response EnterGameResponse) {
	response.Quiz = match
	response.Roster = TableRoster(teams, teamPlayers, players)
	response.PlayerTeamId = playerTeamId
	response.Snapshot = snapshot
	return
}

func InitStartGameResponse(quizId string, teams []Team, teamPlayers []TeamPlayer, players []Player, question Question, snapshot Snapshot) (response StartGameResponse) {
	response.QuizId = quizId
	response.Roster = TableRoster(teams, teamPlayers, players)
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
