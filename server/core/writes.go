package core

import (
	"github.com/atulanand206/go-mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DB struct{}

func (db DB) PlayersCollections() []string {
	return []string{TeamCollection, SubscriberCollection}
}

func (db DB) QuestionsCollections() []string {
	return []string{IndexCollection, QuestionCollection, AnswerCollection}
}

func (db DB) CreateCollections() (err error) {
	err = mongo.CreateCollections(Database, db.QuestionsCollections())
	return
}

func (db DB) DropCollections() (err error) {
	err = mongo.DropCollections(Database, db.PlayersCollections())
	return
}

type Schemas struct{}

func (schemas Schemas) Subscriber() (jsonSchema bson.M) {
	return bson.M{
		"bsonType": "object",
		"required": []string{"tag, playerId, role"},
		"properties": bson.M{
			"tag": bson.M{
				"bsonType":    "string",
				"description": "subscriber must have a tag assigned.",
			},
			"playerId": bson.M{
				"bsonType":  "string",
				"describer": "subscriber must have a valid player id.",
			},
			"role": bson.M{
				"bsonType":  "string",
				"describer": "subscriber must have a role assigned.",
			},
		},
	}
}

func (db DB) CreateSubscriberCollection() (err error) {
	err = mongo.CreateCollection(Database, SubscriberCollection,
		options.CreateCollection().SetValidator(bson.M{
			"$jsonSchema": Schemas{}.Subscriber(),
		}))
	return
}

func (db DB) Init() (err error) {
	if err = db.DropCollections(); err != nil {
		return
	}
	if err = db.CreateSubscriberCollection(); err != nil {
		return
	}

	return
}

func (db DB) Create(request interface{}, collection string) (err error) {
	requestDto, err := mongo.Document(request)
	if err != nil {
		return
	}
	_, err = mongo.Write(Database, collection, *requestDto)
	return
}

func SeedIndexes(indexes []Index) (err error) {
	var indexesDto []interface{}
	for _, t := range indexes {
		indexesDto = append(indexesDto, t)
	}
	_, err = mongo.WriteMany(Database, IndexCollection, indexesDto)
	return err
}

func CreateQuestion(question Question) (err error) {
	requestDto, err := mongo.Document(question)
	if err != nil {
		return
	}
	_, err = mongo.Write(Database, QuestionCollection, *requestDto)
	return
}

func SeedQuestions(questions []Question) (err error) {
	var questionsDto []interface{}
	for _, t := range questions {
		questionsDto = append(questionsDto, t)
	}
	_, err = mongo.WriteMany(Database, QuestionCollection, questionsDto)
	return err
}

func CreateAnswer(answer Answer) (err error) {
	requestDto, err := mongo.Document(answer)
	if err != nil {
		return
	}
	_, err = mongo.Write(Database, AnswerCollection, *requestDto)
	return
}

func SeedAnswers(answers []Answer) (err error) {
	var answersDto []interface{}
	for _, answer := range answers {
		answersDto = append(answersDto, answer)
	}
	_, err = mongo.WriteMany(Database, AnswerCollection, answersDto)
	return err
}

func CreateSnapshot(snapshot Snapshot) (err error) {
	requestDto, err := mongo.Document(snapshot)
	if err != nil {
		return
	}
	_, err = mongo.Write(Database, SnapshotCollection, *requestDto)
	return
}

func CreatePlayer(player Player) (err error) {
	requestDto, err := mongo.Document(player)
	if err != nil {
		return
	}
	_, err = mongo.Write(Database, PlayerCollection, *requestDto)
	return
}

func CreateMatch(match Game) (err error) {
	requestDto, err := mongo.Document(match)
	if err != nil {
		return
	}
	_, err = mongo.Write(Database, MatchCollection, *requestDto)
	return
}

func CreateTeams(teams []Team) (err error) {
	var teamsDto []interface{}
	for _, t := range teams {
		teamsDto = append(teamsDto, t)
	}
	_, err = mongo.WriteMany(Database, TeamCollection, teamsDto)
	return
}

func CreateTeamPlayer(team TeamPlayerRequest) (err error) {
	requestDto, err := mongo.Document(team)
	if err != nil {
		return
	}
	_, err = mongo.Write(Database, TeamPlayerCollection, *requestDto)
	return
}

func (db DB) CreateSubscriber(subscriber Subscriber) (err error) {
	err = db.Create(subscriber, SubscriberCollection)
	return
}

func UpdateMatchQuestions(match Game, question Question) (err error) {
	match.Tags = append(match.Tags, question.Tag)
	match.Active = true
	err = UpdateMatch(match)
	return
}

func UpdateMatch(match Game) (err error) {
	requestDto, err := mongo.Document(match)
	if err != nil {
		return
	}
	_, err = mongo.Update(Database, MatchCollection, bson.M{"_id": match.Id}, bson.D{primitive.E{Key: "$set", Value: *requestDto}})
	return
}

func DeleteTeamPlayers(ids []string) (err error) {
	_, err = mongo.Delete(Database, TeamPlayerCollection, bson.M{"_id": bson.M{"$in": ids}})
	return
}

func (db DB) DeleteSubscribers(tag string, playerIds []string) (err error) {
	_, err = mongo.Delete(Database, TeamPlayerCollection, bson.M{
		"tag":      tag,
		"playerId": bson.M{"$in": playerIds},
		"active":   true})
	return
}

func (db DB) DeleteSubscriber(playerId string) (err error) {
	_, err = mongo.Delete(Database, TeamPlayerCollection, bson.M{
		"playerId": playerId,
		"active":   true})
	return
}
