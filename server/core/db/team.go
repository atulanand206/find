package db

import (
	"github.com/atulanand206/find/server/core/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type TeamCrud struct {
	Db DBConn
}

func (crud TeamCrud) CreateTeams(teams []models.Team) (err error) {
	var teamsDto []interface{}
	for _, t := range teams {
		teamsDto = append(teamsDto, t)
	}
	return crud.Db.CreateMany(teamsDto, TeamCollection)
}

func (crud TeamCrud) FindTeams(matchId string) (teams []models.Team, err error) {
	findOptions := &options.FindOptions{}
	sort := bson.D{}
	sort = append(sort, bson.E{Key: "name", Value: 1})
	findOptions.SetSort(sort)
	cursor, err := crud.Db.Find(TeamCollection,
		bson.M{"quiz_id": matchId}, findOptions)

	if err != nil {
		return
	}
	teams, err = DecodeTeams(cursor)
	return
}

func (crud TeamCrud) UpdateTeam(team models.Team) (err error) {
	_, err = crud.Db.Update(TeamCollection, bson.M{"_id": team.Id}, team)
	return
}
