package core

import (
	"fmt"

	"github.com/atulanand206/go-mongo"
	"github.com/xorcare/pointer"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func FindMatch(matchId string) (match Game, err error) {
	dto := mongo.FindOne(Database, MatchCollection, bson.M{"_id": matchId}, &options.FindOneOptions{})
	if err = dto.Err(); err != nil {
		return
	}
	match, err = DecodeMatch(dto)
	return
}

func FindQuestion(questionId string) (question Question, err error) {
	dto := mongo.FindOne(Database, MatchCollection, bson.M{"_id": questionId}, &options.FindOneOptions{})
	if err = dto.Err(); err != nil {
		return
	}
	question, err = DecodeQuestion(dto)
	return
}

func FindPlayer(emailId string) (player Player, err error) {
	dto := mongo.FindOne(Database, PlayerCollection, bson.M{"email": emailId}, &options.FindOneOptions{})
	if err = dto.Err(); err != nil {
		return
	}
	player, err = DecodePlayer(dto)
	return
}

func FindIndexForTag(tag string) (index Index, err error) {
	dto := mongo.FindOne(Database, IndexCollection, bson.M{"tag": tag}, &options.FindOneOptions{})
	if err = dto.Err(); err != nil {
		return
	}
	index, err = DecodeIndex(dto)
	return
}

func FindAnswer(questionId string) (answer Answer, err error) {
	dto := mongo.FindOne(Database, AnswerCollection, bson.M{"question_id": questionId}, &options.FindOneOptions{})
	if err = dto.Err(); err != nil {
		return
	}
	answer, err = DecodeAnswer(dto)
	return
}

func FindIndex() (indexes []Index, err error) {
	cursor, err := mongo.Find(Database, IndexCollection, bson.M{}, &options.FindOptions{})
	if err != nil {
		return
	}
	indexes, err = DecodeIndexes(cursor)
	return
}

func FindQuestionsForIndex(index Index, limit int64) (questions []Question, err error) {
	cursor, err := mongo.Find(Database, QuestionCollection,
		bson.M{"tag": index.Id}, &options.FindOptions{Limit: pointer.Int64(limit)})
	if err != nil {
		return
	}
	questions, err = DecodeQuestions(cursor)
	return
}

func FindQuestionsFromIndexes(indexes []Index, limit int64) (questions []Question, err error) {
	questions = make([]Question, 0)
	for _, indx := range indexes {
		indxQues, er := FindQuestionsForIndex(indx, limit)
		if er != nil {
			return
		}
		questions = append(questions, indxQues...)
	}
	return
}

func FindTeams(match Game) (teams []Team, err error) {
	findOptions := &options.FindOptions{}
	sort := bson.D{}
	sort = append(sort, bson.E{Key: "name", Value: -1})
	findOptions.SetSort(sort)
	cursor, err := mongo.Find(Database, TeamCollection,
		bson.M{"quiz_id": match.Id}, findOptions)

	if err != nil {
		return
	}
	teams, err = DecodeTeams(cursor)
	fmt.Println(cursor)
	fmt.Println(teams)
	fmt.Println(err)
	return
}

func FindPlayers(teamPlayers []TeamPlayer) (players []Player, err error) {
	playerIds := make([]string, 0)
	for _, v := range teamPlayers {
		playerIds = append(playerIds, v.PlayerId)
	}
	findOptions := &options.FindOptions{}
	cursor, err := mongo.Find(Database, PlayerCollection,
		bson.M{"_id": bson.M{"$in": playerIds}}, findOptions)
	if err != nil {
		return
	}
	players, err = DecodePlayers(cursor)
	return
}

func FindTeamPlayers(teams []Team) (teamPlayers []TeamPlayer, err error) {
	teamIds := make([]string, 0)
	for _, v := range teams {
		teamIds = append(teamIds, v.Id)
	}
	findOptions := &options.FindOptions{}
	sort := bson.D{}
	sort = append(sort, bson.E{Key: "team_id", Value: -1})
	findOptions.SetSort(sort)
	cursor, err := mongo.Find(Database, TeamPlayerCollection,
		bson.M{"team_id": bson.M{"$in": teamIds}}, findOptions)
	if err != nil {
		return
	}
	teamPlayers, err = DecodeTeamPlayers(cursor)
	return
}

func FindSnapshots(matchId string) (snapshots []Snapshot, err error) {
	cursor, err := mongo.Find(Database, SnapshotCollection,
		bson.M{"quiz_id": matchId}, &options.FindOptions{})
	if err != nil {
		return
	}
	snapshots, err = DecodeSnapshots(cursor)
	return
}

func FindLatestSnapshot(matchId string) (snapshot Snapshot, err error) {
	findOptions := &options.FindOneOptions{}
	sort := bson.D{}
	sort = append(sort, bson.E{Key: "timestamp", Value: -1})
	findOptions.SetSort(sort)
	dto := mongo.FindOne(Database, SnapshotCollection,
		bson.M{"quiz_id": matchId}, findOptions)
	if err = dto.Err(); err != nil {
		return
	}
	snapshot, err = DecodeSnapshot(dto)
	return
}

func FindQuestionSnapshots(matchId string, questionId string) (snapshot []Snapshot, err error) {
	findOptions := &options.FindOptions{}
	sort := bson.D{}
	sort = append(sort, bson.E{Key: "timestamp", Value: -1})
	findOptions.SetSort(sort)
	cursor, err := mongo.Find(Database, SnapshotCollection,
		bson.M{"quiz_id": matchId, "question_id": questionId}, findOptions)
	if err != nil {
		return
	}
	snapshot, err = DecodeSnapshots(cursor)
	return
}
