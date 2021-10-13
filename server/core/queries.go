package core

import (
	"github.com/atulanand206/go-mongo"
	"github.com/xorcare/pointer"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (db DB) FindMatch(matchId string) (match Game, err error) {
	dto := mongo.FindOne(Database, MatchCollection, bson.M{"_id": matchId}, &options.FindOneOptions{})
	if err = dto.Err(); err != nil {
		return
	}
	match, err = DecodeMatch(dto)
	return
}

func (db DB) FindQuestion(questionId string) (question Question, err error) {
	dto := mongo.FindOne(Database, MatchCollection, bson.M{"_id": questionId}, &options.FindOneOptions{})
	if err = dto.Err(); err != nil {
		return
	}
	question, err = DecodeQuestion(dto)
	return
}

func (db DB) FindPlayer(emailId string) (player Player, err error) {
	dto := mongo.FindOne(Database, PlayerCollection, bson.M{"email": emailId}, &options.FindOneOptions{})
	if err = dto.Err(); err != nil {
		return
	}
	player, err = DecodePlayer(dto)
	return
}

func (db DB) FindIndexForTag(tag string) (index Index, err error) {
	dto := mongo.FindOne(Database, IndexCollection, bson.M{"tag": tag}, &options.FindOneOptions{})
	if err = dto.Err(); err != nil {
		return
	}
	index, err = DecodeIndex(dto)
	return
}

func (db DB) FindAnswer(questionId string) (answer Answer, err error) {
	dto := mongo.FindOne(Database, AnswerCollection, bson.M{"question_id": questionId}, &options.FindOneOptions{})
	if err = dto.Err(); err != nil {
		return
	}
	answer, err = DecodeAnswer(dto)
	return
}

func (db DB) FindIndex() (indexes []Index, err error) {
	cursor, err := mongo.Find(Database, IndexCollection, bson.M{}, &options.FindOptions{})
	if err != nil {
		return
	}
	indexes, err = DecodeIndexes(cursor)
	return
}

func (db DB) FindQuestionsForIndex(index Index, limit int64) (questions []Question, err error) {
	cursor, err := mongo.Find(Database, QuestionCollection,
		bson.M{"tag": index.Id}, &options.FindOptions{Limit: pointer.Int64(limit)})
	if err != nil {
		return
	}
	questions, err = DecodeQuestions(cursor)
	return
}

func (db DB) FindQuestionsFromIndexes(indexes []Index, limit int64) (questions []Question, err error) {
	questions = make([]Question, 0)
	for _, indx := range indexes {
		indxQues, er := db.FindQuestionsForIndex(indx, limit)
		if er != nil {
			return
		}
		questions = append(questions, indxQues...)
	}
	return
}

func (db DB) FindActiveMatches() (matches []Game, err error) {
	findOptions := &options.FindOptions{}
	sort := bson.D{}
	sort = append(sort, bson.E{Key: "quizmaster.name", Value: 1})
	findOptions.SetSort(sort)
	cursor, err := mongo.Find(Database, MatchCollection,
		bson.M{"active": true}, findOptions)

	if err != nil {
		return
	}
	matches, err = DecodeMatches(cursor)
	return
}

func (db DB) FindTeams(match Game) (teams []Team, err error) {
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
	return
}

func (db DB) FindTeamsMatches(match []Game) (teams []Team, err error) {
	matchIds := make([]string, 0)
	for _, v := range match {
		matchIds = append(matchIds, v.Id)
	}
	findOptions := &options.FindOptions{}
	sort := bson.D{}
	sort = append(sort, bson.E{Key: "name", Value: -1})
	findOptions.SetSort(sort)
	cursor, err := mongo.Find(Database, TeamCollection,
		bson.M{"quiz_id": bson.M{"$in": matchIds}}, findOptions)

	if err != nil {
		return
	}
	teams, err = DecodeTeams(cursor)
	return
}

func (db DB) FindPlayers(teamPlayers []Subscriber) (players []Player, err error) {
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

func (db DB) FindTPs(teamPlayers []Subscriber) (teams []Team, err error) {
	teamIds := make([]string, 0)
	for _, v := range teamPlayers {
		teamIds = append(teamIds, v.Tag)
	}
	findOptions := &options.FindOptions{}
	cursor, err := mongo.Find(Database, TeamCollection,
		bson.M{"_id": bson.M{"$in": teamIds}}, findOptions)
	if err != nil {
		return
	}
	teams, err = DecodeTeams(cursor)
	return
}

func (db DB) FindActiveTeamMatches(teams []Team) (matches []Game, err error) {
	matchIds := make([]string, 0)
	for _, v := range teams {
		matchIds = append(matchIds, v.QuizId)
	}
	findOptions := &options.FindOptions{}
	cursor, err := mongo.Find(Database, MatchCollection,
		bson.M{"_id": bson.M{"$in": matchIds}, "active": true}, findOptions)
	if err != nil {
		return
	}
	matches, err = DecodeMatches(cursor)
	return
}

func (db DB) FindTeamPlayers(teams []Team) (teamPlayers []Subscriber, err error) {
	teamIds := make([]string, 0)
	for _, v := range teams {
		teamIds = append(teamIds, v.Id)
	}
	findOptions := &options.FindOptions{}
	sort := bson.D{}
	sort = append(sort, bson.E{Key: "tag", Value: 1})
	findOptions.SetSort(sort)
	cursor, err := mongo.Find(Database, SubscriberCollection,
		bson.M{"tag": bson.M{"$in": teamIds}}, findOptions)
	if err != nil {
		return
	}
	teamPlayers, err = DecodeTeamPlayers(cursor)
	return
}

func (db DB) FindIds(teamPlayers []Subscriber) (playerIds []string) {
	playerIds = make([]string, 0)
	for _, v := range teamPlayers {
		playerIds = append(playerIds, v.PlayerId)
	}
	return
}

func (db DB) FindPlayerTeams(playerId string) (teamPlayers []Subscriber, err error) {
	findOptions := &options.FindOptions{}
	cursor, err := mongo.Find(Database, SubscriberCollection,
		bson.M{"player_id": bson.M{"$in": playerId}}, findOptions)
	if err != nil {
		return
	}
	teamPlayers, err = DecodeTeamPlayers(cursor)
	return
}

func (db DB) FindSnapshots(matchId string) (snapshots []Snapshot, err error) {
	cursor, err := mongo.Find(Database, SnapshotCollection,
		bson.M{"quiz_id": matchId}, &options.FindOptions{})
	if err != nil {
		return
	}
	snapshots, err = DecodeSnapshots(cursor)
	return
}

func (db DB) FindLatestSnapshot(matchId string) (snapshot Snapshot, err error) {
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

func (db DB) FindQuestionSnapshots(matchId string, questionId string) (snapshot []Snapshot, err error) {
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

func (db DB) FindSubscribers(tag string, role Role) (subscribers []Subscriber, err error) {
	findOptions := &options.FindOptions{}
	sort := bson.D{}
	findOptions.SetSort(sort)
	cursor, err := mongo.Find(Database, SubscriberCollection,
		bson.M{"tag": tag, "role": role.String()}, findOptions)
	if err != nil {
		return
	}
	subscribers, err = DecodeSubscribers(cursor)
	return
}

func (db DB) FindSubscribersForTag(tags []string) (subscribers []Subscriber, err error) {
	findOptions := &options.FindOptions{}
	sort := bson.D{}
	findOptions.SetSort(sort)
	cursor, err := mongo.Find(Database, SubscriberCollection,
		bson.M{"tag": bson.M{"$in": tags}}, findOptions)
	if err != nil {
		return
	}
	subscribers, err = DecodeSubscribers(cursor)
	return
}

func (db DB) FindSubscriberForTagAndPlayerId(tag string, playerId string) (subscriber Subscriber, err error) {
	dto := mongo.FindOne(Database, SubscriberCollection,
		bson.M{"tag": tag, "player_id": playerId}, &options.FindOneOptions{})
	if err = dto.Err(); err != nil {
		return
	}
	subscriber, err = DecodeSubscriber(dto)
	return
}

func (db DB) FindSubscriptionsForPlayerId(playerId string) (subscribers []Subscriber, err error) {
	findOptions := &options.FindOptions{}
	sort := bson.D{}
	findOptions.SetSort(sort)
	cursor, err := mongo.Find(Database, SubscriberCollection,
		bson.M{"player_id": playerId}, findOptions)
	if err != nil {
		return
	}
	subscribers, err = DecodeSubscribers(cursor)
	return
}
