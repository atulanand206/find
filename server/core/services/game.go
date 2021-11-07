package services

import (
	e "errors"

	"github.com/atulanand206/find/server/core/actions"
	"github.com/atulanand206/find/server/core/db"
	"github.com/atulanand206/find/server/core/errors"
	"github.com/atulanand206/find/server/core/models"
	"github.com/atulanand206/find/server/core/utils"
)

type Creators struct {
	InstanceCreator models.Creator
	MessageCreator  models.WebsocketMessageCreator
	ErrorCreator    errors.ErrorMessageCreator
}

type Service struct {
	AuthService       AuthService
	PermissionService PermissionService
	MatchService      MatchService
	SubscriberService SubscriberService
	PlayerService     PlayerService
	TeamService       TeamService
	QuestionService   QuestionService
	SnapshotService   SnapshotService
	Validator         Validator
	TargetService     TargetService
	Creators          Creators
}

func Init(Db db.DB) (service Service) {
	service = Service{}

	creators := Creators{}
	creators.InstanceCreator = models.Creator{}
	creators.MessageCreator = models.WebsocketMessageCreator{}
	creators.ErrorCreator = errors.ErrorMessageCreator{}
	service.Creators = creators
	service.Validator = Validator{}
	service.AuthService = AuthService{}
	service.PermissionService = PermissionService{}

	service.TargetService = TargetService{}
	service.SubscriberService = SubscriberService{Crud: db.SubscriberCrud{Db: Db}, TargetService: service.TargetService, Creators: service.Creators}
	service.MatchService = MatchService{Crud: db.MatchCrud{Db: Db}, SubscriberService: service.SubscriberService}
	service.TeamService = TeamService{Crud: db.TeamCrud{Db: Db}, SubscriberService: service.SubscriberService}
	service.PlayerService = PlayerService{Crud: db.PlayerCrud{Db: Db}}
	service.SnapshotService = SnapshotService{Crud: db.SnapshotCrud{Db: Db}}
	service.QuestionService = QuestionService{Crud: db.QuestionCrud{Db: Db}}

	return service
}

func (service Service) GenerateGameResponse(request models.Request) (response models.Snapshot, err error) {

	return
}

func (service Service) GenerateBeginGameResponse(player models.Player) (res models.LoginResponse, err error) {
	response, err := service.PlayerService.FindOrCreatePlayer(player)
	if err != nil {
		return
	}

	tokens, err := service.AuthService.GenerateTokens(response)
	if err != nil {
		return
	}

	quizmaster := service.PermissionService.HasPermissions(player.Id)
	res = models.LoginResponse{Player: response, Tokens: tokens, Quizmaster: quizmaster}
	return
}

func (service Service) GenerateCreateGameResponse(quizmaster models.Player, specs models.Specs) (response models.GameResponse, err error) {
	player := quizmaster
	player, err = service.PlayerService.FindOrCreatePlayer(player)
	if err != nil {
		return
	}

	quiz, err := service.MatchService.CreateMatch(player, specs)
	if err != nil {
		return
	}

	teams, err := service.TeamService.CreateTeams(quiz)
	if err != nil {
		return
	}

	snapshot, err := service.SnapshotService.InitialSnapshot(quiz.Id, teams)
	if err != nil {
		return
	}

	return service.SubscriberService.SubscribeAndRespond(quiz, player, snapshot, actions.QUIZMASTER)
}

func (service Service) GenerateEnterGameResponse(request models.Request) (response models.GameResponse, err error) {
	player, err := service.PlayerService.FindPlayerByEmail(request.Person.Email)
	if err != nil {
		return
	}

	match, teams, teamPlayers, _, _, snapshot, err := service.FindMatchFull(request.QuizId)
	if err != nil {
		return
	}

	if utils.IsQuizMasterInMatch(match, player) {
		match, _, _, _, roster, snapshot, er := service.FindMatchFull(request.QuizId)
		if er != nil {
			err = er
			return
		}

		snapshot, er = service.SnapshotService.SnapshotJoin(snapshot, roster)
		if er != nil {
			err = er
			return
		}

		return service.SubscriberService.SubscribeAndRespond(match, player, snapshot, actions.QUIZMASTER)
	}

	if utils.IsPlayerInTeams(teamPlayers, player) {
		match, _, _, _, roster, snapshot, er := service.FindMatchFull(request.QuizId)
		if er != nil {
			err = er
			return
		}

		snapshot, er = service.SnapshotService.SnapshotJoin(snapshot, roster)
		if er != nil {
			err = er
			return
		}
		return service.SubscriberService.SubscribeAndRespond(match, player, snapshot, actions.PLAYER)
	}

	_, err = service.TeamService.FindAndFillTeamVacancy(match, teams, player)
	if err != nil {
		return
	}

	match, _, _, _, roster, snapshot, err := service.FindMatchFull(request.QuizId)
	if err != nil {
		return
	}

	snapshot, err = service.SnapshotService.SnapshotJoin(snapshot, roster)
	if err != nil {
		return
	}

	return service.SubscriberService.SubscribeAndRespond(match, player, snapshot, actions.PLAYER)
}

func (service Service) GenerateFullMatchResponse(quizId string) (response models.Snapshot, err error) {
	_, _, _, _, _, snapshot, err := service.FindMatchFull(quizId)
	if err != nil {
		return
	}

	response = snapshot
	return
}

func (service Service) GenerateWatchGameResponse(request models.Request) (response models.GameResponse, err error) {
	audience, err := service.PlayerService.FindPlayerByEmail(request.Person.Email)
	if err != nil {
		return
	}

	match, _, _, _, _, snapshot, err := service.FindMatchFull(request.QuizId)
	if err != nil {
		return
	}

	return service.SubscriberService.SubscribeAndRespond(match, audience, snapshot, actions.AUDIENCE)
}

func (service Service) GenerateStartGameResponse(request models.Request) (response models.Snapshot, err error) {
	match, teams, teamPlayers, _, roster, snapshot, err := service.FindMatchFull(request.QuizId)
	if err != nil {
		return
	}

	if result := utils.MatchFull(match, teamPlayers); !result {
		err = e.New(errors.Err_WaitingForPlayers)
		return
	}

	question, err := service.QuestionService.FindQuestionForMatch(match)
	if err != nil {
		return
	}

	match.Started = true
	if _, err = service.MatchService.Crud.UpdateMatchQuestions(match, question); err != nil {
		err = e.New(errors.Err_MatchNotUpdated)
		return
	}

	snapshot, err = service.SnapshotService.SnapshotStart(snapshot, roster, question, teams[0].Id)
	if err != nil {
		return
	}

	response = snapshot
	return
}

func (service Service) GenerateQuestionHintResponse(request models.Request) (response models.Snapshot, err error) {
	_, _, _, _, roster, snapshot, err := service.FindMatchFull(request.QuizId)
	if err != nil {
		return
	}

	answer, err := service.QuestionService.Crud.FindAnswer(request.QuestionId)
	if err != nil {
		err = e.New(errors.Err_QuestionNotPresent)
		return
	}

	snapshot, err = service.SnapshotService.SnapshotHint(snapshot, roster, answer.Hint)
	if err != nil {
		return
	}

	response = snapshot
	return
}

func (service Service) GenerateQuestionAnswerResponse(request models.Request) (response models.Snapshot, err error) {
	match, teams, _, _, _, snapshot, err := service.FindMatchFull(request.QuizId)
	if err != nil {
		return
	}

	teamId := snapshot.TeamSTurn
	var team models.Team
	for _, tm := range teams {
		if tm.Id == teamId {
			team = tm
		}
	}
	points := match.Specs.Points
	team.Score += models.ScoreAnswer(points, snapshot.RoundNo)

	err = service.TeamService.Crud.UpdateTeam(team)
	if err != nil {
		err = e.New(errors.Err_TeamNotUpdated)
		return
	}

	answer, err := service.QuestionService.Crud.FindAnswer(request.QuestionId)
	if err != nil {
		err = e.New(errors.Err_QuestionNotPresent)
		return
	}

	_, _, _, _, roster, _, err := service.FindMatchFull(request.QuizId)
	if err != nil {
		return
	}

	snapshot, err = service.SnapshotService.SnapshotAnswer(snapshot, roster, answer, points)
	if err != nil {
		return
	}

	response = snapshot
	return
}

func (service Service) GenerateNextQuestionResponse(request models.Request) (response models.Snapshot, err error) {

	match, _, _, _, roster, snapshot, err := service.FindMatchFull(request.QuizId)
	if err != nil {
		return
	}

	if !utils.QuestionCanBeAdded(match) {
		err = e.New(errors.Err_QuestionsNotLeft)
		return
	}

	question, err := service.QuestionService.FindQuestionForMatch(match)
	if err != nil {
		return
	}

	if _, err = service.MatchService.Crud.UpdateMatchQuestions(match, question); err != nil {
		err = e.New(errors.Err_MatchNotUpdated)
		return
	}

	snapshot, err = service.SnapshotService.SnapshotNext(snapshot, roster, question, snapshot.TeamSTurn)
	if err != nil {
		return
	}

	response = snapshot
	return
}

func (service Service) GeneratePassQuestionResponse(request models.Request) (response models.Snapshot, err error) {
	match, _, _, _, roster, snapshot, err := service.FindMatchFull(request.QuizId)
	if err != nil {
		return
	}

	snapshot, err = service.SnapshotService.SnapshotPass(snapshot, roster, snapshot.TeamSTurn, snapshot.RoundNo, match.Specs)
	if err != nil {
		return
	}

	response = snapshot
	return
}

func (service Service) GenerateScoreResponse(request models.Request) (response models.ScoreResponse, err error) {
	snapshots, err := service.SnapshotService.Crud.FindSnapshotsForMatch(request.QuizId)
	if err != nil {
		err = e.New(errors.Err_SnapshotNotPresent)
		return
	}

	response = models.InitScoreResponse(request.QuizId, snapshots)
	return
}

func (service Service) DeletePlayerLiveSession(playerId string) (res models.WebsocketMessage, targets map[string]bool, err error) {
	subscribers, err := service.SubscriberService.Crud.FindSubscriptionsForPlayerId(playerId)
	if err != nil {
		err = e.New(err.Error())
		return
	}

	tags := make([]string, 0)
	for _, subscriber := range subscribers {
		tags = append(tags, subscriber.Tag)
	}

	subscribers, err = service.SubscriberService.FindSubscribersForTag(tags)
	if err != nil {
		err = e.New(err.Error())
		return
	}

	targets = make(map[string]bool)
	for _, subscriber := range subscribers {
		if playerId != subscriber.PlayerId {
			targets[subscriber.PlayerId] = true
		}
	}

	res = service.Creators.MessageCreator.InitWebSocketMessage(actions.S_REFRESH, "Player dropped. Please refresh.")
	err = service.SubscriberService.Crud.DeleteSubscriber(playerId)
	return
}
