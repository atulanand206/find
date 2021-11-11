package services

import (
	e "errors"

	"github.com/atulanand206/find/server/core/actions"
	"github.com/atulanand206/find/server/core/db"
	"github.com/atulanand206/find/server/core/errors"
	"github.com/atulanand206/find/server/core/models"
	"github.com/atulanand206/find/server/core/utils"
	"go.mongodb.org/mongo-driver/bson"
)

type SnapshotService struct {
	Crud db.SnapshotCrud
}

func (service SnapshotService) CreateSnapshot(snapshot models.Snapshot) (err error) {
	if err = service.Crud.CreateSnapshot(snapshot); err != nil {
		err = e.New(errors.Err_SnapshotNotCreated)
		return
	}
	return
}

func (service SnapshotService) FindSnapshotsForMatch(matchId string) (snapshots []models.Snapshot, err error) {
	return service.Crud.FindSnapshots(bson.M{"quiz_id": matchId})
}

func (service SnapshotService) FindSnapshotsForQuestion(quizId string, questionId string, eventType string) (snapshots []models.Snapshot, err error) {
	return service.Crud.FindSnapshots(bson.M{"quiz_id": quizId, "question_id": questionId, "event_type": eventType})
}

func (service SnapshotService) InitialSnapshot(quizId string, teams []models.Team) (response models.Snapshot, err error) {
	roster := utils.TableRoster(teams, []models.Subscriber{}, []models.Player{})
	snapshot := models.InitialSnapshot(quizId, roster)
	err = service.CreateSnapshot(snapshot)
	if err != nil {
		return
	}
	response = snapshot
	return
}

func (service SnapshotService) SnapshotJoin(snapshot models.Snapshot, roster []models.TeamRoster) (response models.Snapshot, err error) {
	snapshot = models.SnapshotWithJoinPlayer(snapshot, roster)
	err = service.CreateSnapshot(snapshot)
	if err != nil {
		return
	}
	response = snapshot
	return
}

func (service SnapshotService) SnapshotDrop(snapshot models.Snapshot, roster []models.TeamRoster) (response models.Snapshot, err error) {
	snapshot = models.SnapshotWithDropPlayer(snapshot, roster)
	err = service.CreateSnapshot(snapshot)
	if err != nil {
		return
	}
	response = snapshot
	return
}

func (service SnapshotService) SnapshotStart(snapshot models.Snapshot, roster []models.TeamRoster, question models.Question, teamsTurn string) (response models.Snapshot, err error) {
	snapshot = models.SnapshotWithStart(snapshot, roster, question, teamsTurn)
	err = service.CreateSnapshot(snapshot)
	if err != nil {
		return
	}
	response = snapshot
	return
}

func (service SnapshotService) SnapshotHint(snapshot models.Snapshot, roster []models.TeamRoster, hint []string) (response models.Snapshot, err error) {
	snapshot = models.SnapshotWithHint(snapshot, hint, roster)
	err = service.CreateSnapshot(snapshot)
	if err != nil {
		return
	}
	response = snapshot
	return
}

func (service SnapshotService) SnapshotAnswer(snapshot models.Snapshot, roster []models.TeamRoster, answer models.Answer, points int) (response models.Snapshot, err error) {
	snapshot = models.SnapshotWithAnswer(snapshot, answer.Answer, points, roster)
	err = service.CreateSnapshot(snapshot)
	if err != nil {
		return
	}
	response = snapshot
	return
}

func (service SnapshotService) SnapshotNext(snapshot models.Snapshot, roster []models.TeamRoster, question models.Question, teamsTurn string) (response models.Snapshot, err error) {
	teamsTurn = utils.NextTeam(roster, teamsTurn)
	snapshot = models.SnapshotWithNext(snapshot, roster, teamsTurn, question)
	err = service.CreateSnapshot(snapshot)
	if err != nil {
		return
	}
	response = snapshot
	return
}

func (service SnapshotService) SnapshotPass(snapshot models.Snapshot, roster []models.TeamRoster, teamsTurn string, roundNo int, specs models.Specs) (response models.Snapshot, err error) {
	snapshots, err := service.FindSnapshotsForQuestion(snapshot.QuizId, snapshot.QuestionId, actions.PASS.String())
	if err != nil {
		return
	}

	if len(snapshots) > 0 {
		if roundNo == specs.Rounds && len(snapshots)%specs.Teams == 0 {
			snapshot.CanPass = false
		}
		if (len(snapshots)+1)%specs.Teams == 0 {
			teamsTurn = utils.NextTeam(roster, teamsTurn)
			roundNo++
		}
	} else {
		teamsTurn = utils.NextTeam(roster, teamsTurn)
	}
	snapshot = models.SnapshotWithPass(snapshot, roster, teamsTurn, roundNo)
	err = service.CreateSnapshot(snapshot)
	if err != nil {
		return
	}
	response = snapshot
	return
}
