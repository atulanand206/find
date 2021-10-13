package core

import (
	"fmt"
	"math"
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

func (creator Creator) InitSubscriber(tag string, player Player, role string) (subscriber Subscriber) {
	subscriber.Tag = tag
	subscriber.PlayerId = player.Id
	subscriber.Role = role
	subscriber.Active = true
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

func (creator Creator) InitEnterGameResponse(match Game, roster []TeamRoster, snapshot Snapshot) (response EnterGameResponse) {
	response.Quiz = match
	response.Roster = roster
	response.Snapshot = snapshot
	return
}

func InitStartGameResponse(quizId string, roster []TeamRoster, snapshot Snapshot) (response StartGameResponse) {
	response.QuizId = quizId
	response.Roster = roster
	response.Snapshot = snapshot
	return
}

func InitScoreResponse(request ScoreRequest, snapshots []Snapshot) (response ScoreResponse) {
	response.QuizId = request.QuizId
	response.Snapshots = snapshots
	return
}

func InitialSnapshot(quizId string, question Question, teams []TeamRoster, teamsTurn string) (response Snapshot) {
	response.QuizId = quizId
	response.QuestionId = question.Id
	response.Roster = teams
	response.TeamSTurn = teamsTurn
	response.EventType = START.String()
	response.Question = question.Statements
	response.Score = 0
	response.QuestionNo = 1
	response.RoundNo = 1
	response.Timestamp = time.Now().String()
	return
}

func SnapshotWithAnswer(snapshot Snapshot, answer []string, matchPoints int, teams []TeamRoster) (response Snapshot) {
	snapshot.Answer = answer
	snapshot.EventType = RIGHT.String()
	snapshot.Score = ScoreAnswer(matchPoints, snapshot.RoundNo)
	snapshot.Roster = teams
	response = snapshot
	return
}

func ScoreAnswer(matchPoints int, roundNo int) int {
	return matchPoints / int(math.Pow(2, float64(roundNo)))
}

func SnapshotWithHint(snapshot Snapshot, hint []string, teams []TeamRoster) (response Snapshot) {
	snapshot.Hint = hint
	snapshot.Roster = teams
	snapshot.EventType = HINT.String()
	snapshot.Score = 0
	response = snapshot
	return
}

func SnapshotWithPass(snapshot Snapshot, teams []TeamRoster, team_s_turn string) (response Snapshot) {
	snapshot.TeamSTurn = team_s_turn
	snapshot.Roster = teams
	snapshot.EventType = PASS.String()
	snapshot.Score = 0
	response = snapshot
	return
}

func SnapshotWithNext(snapshot Snapshot, teams []TeamRoster, team_s_turn string, question Question) (response Snapshot) {
	snapshot.TeamSTurn = team_s_turn
	snapshot.Roster = teams
	snapshot.EventType = NEXT.String()
	snapshot.Score = 0
	snapshot.RoundNo = 1
	snapshot.Question = question.Statements
	snapshot.QuestionNo = snapshot.QuestionNo + 1
	snapshot.QuestionId = question.Id
	response = snapshot
	return
}

func (creator WebsocketMessageCreator) InitWebSocketMessageFailure() (response WebsocketMessage) {
	response = creator.InitWebSocketMessage(Failure, Err_SocketRequestFailed)
	return
}

func (creator WebsocketMessageCreator) InitWebSocketMessage(action Action, content string) (response WebsocketMessage) {
	response.Action = action.String()
	response.Content = content
	return
}
