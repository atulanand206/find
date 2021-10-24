package core

import "errors"

type SnapshotService struct {
	db DB
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

func (service SnapshotService) SnapshotPass(snapshot Snapshot, roster []TeamRoster, teamsTurn string) (response Snapshot, err error) {
	teamsTurn = NextTeam(roster, teamsTurn)
	snapshot = SnapshotWithPass(snapshot, roster, teamsTurn)
	err = service.CreateSnapshot(snapshot)
	if err != nil {
		return
	}
	response = snapshot
	return
}

func (service SnapshotService) CreateSnapshot(snapshot Snapshot) (err error) {
	if err = service.db.CreateSnapshot(snapshot); err != nil {
		err = errors.New(Err_SnapshotNotCreated)
		return
	}
	return
}
