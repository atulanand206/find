package core

import (
	"errors"
)

type SnapshotService struct {
	crud SnapshotCrud
}

func (service SnapshotService) InitialSnapshot(quizId string, teams []Team) (response Snapshot, err error) {
	roster := TableRoster(teams, []Subscriber{}, []Player{})
	snapshot := InitialSnapshot(quizId, roster)
	err = service.CreateSnapshot(snapshot)
	if err != nil {
		return
	}
	response = snapshot
	return
}

func (service SnapshotService) SnapshotJoin(snapshot Snapshot, roster []TeamRoster) (response Snapshot, err error) {
	snapshot = SnapshotWithJoinPlayer(snapshot, roster)
	err = service.CreateSnapshot(snapshot)
	if err != nil {
		return
	}
	response = snapshot
	return
}

func (service SnapshotService) SnapshotDrop(snapshot Snapshot, roster []TeamRoster) (response Snapshot, err error) {
	snapshot = SnapshotWithDropPlayer(snapshot, roster)
	err = service.CreateSnapshot(snapshot)
	if err != nil {
		return
	}
	response = snapshot
	return
}

func (service SnapshotService) SnapshotStart(snapshot Snapshot, roster []TeamRoster, question Question, teamsTurn string) (response Snapshot, err error) {
	snapshot = SnapshotWithStart(snapshot, roster, question, teamsTurn)
	err = service.CreateSnapshot(snapshot)
	if err != nil {
		return
	}
	response = snapshot
	return
}

func (service SnapshotService) SnapshotHint(snapshot Snapshot, roster []TeamRoster, hint []string) (response Snapshot, err error) {
	snapshot = SnapshotWithHint(snapshot, hint, roster)
	err = service.CreateSnapshot(snapshot)
	if err != nil {
		return
	}
	response = snapshot
	return
}

func (service SnapshotService) SnapshotAnswer(snapshot Snapshot, roster []TeamRoster, answer Answer, points int) (response Snapshot, err error) {
	snapshot = SnapshotWithAnswer(snapshot, answer.Answer, points, roster)
	err = service.CreateSnapshot(snapshot)
	if err != nil {
		return
	}
	response = snapshot
	return
}

func (service SnapshotService) SnapshotNext(snapshot Snapshot, roster []TeamRoster, question Question, teamsTurn string) (response Snapshot, err error) {
	teamsTurn = NextTeam(roster, teamsTurn)
	snapshot = SnapshotWithNext(snapshot, roster, teamsTurn, question)
	err = service.CreateSnapshot(snapshot)
	if err != nil {
		return
	}
	response = snapshot
	return
}

func (service SnapshotService) SnapshotPass(snapshot Snapshot, roster []TeamRoster, teamsTurn string, roundNo int, specs Specs) (response Snapshot, err error) {
	snapshots, err := service.crud.FindSnapshotsForQuestion(snapshot.QuizId, snapshot.QuestionId, PASS.String())
	if err != nil {
		return
	}

	if len(snapshots) > 0 {
		if roundNo == specs.Rounds && len(snapshots)%specs.Teams == 0 {
			snapshot.CanPass = false
		}
		if (len(snapshots)+1)%specs.Teams == 0 {
			teamsTurn = NextTeam(roster, teamsTurn)
			roundNo++
		}
	} else {
		teamsTurn = NextTeam(roster, teamsTurn)
	}
	snapshot = SnapshotWithPass(snapshot, roster, teamsTurn, roundNo)
	err = service.CreateSnapshot(snapshot)
	if err != nil {
		return
	}
	response = snapshot
	return
}

func (service SnapshotService) CreateSnapshot(snapshot Snapshot) (err error) {
	if err = service.crud.CreateSnapshot(snapshot); err != nil {
		err = errors.New(Err_SnapshotNotCreated)
		return
	}
	return
}
