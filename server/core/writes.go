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
			"player_id": bson.M{
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

func (db DB) SeedIndexes(indexes []Index) (err error) {
	var indexesDto []interface{}
	for _, t := range indexes {
		indexesDto = append(indexesDto, t)
	}
	_, err = mongo.WriteMany(Database, IndexCollection, indexesDto)
	return err
}

func (db DB) CreateQuestion(question Question) (err error) {
	requestDto, err := mongo.Document(question)
	if err != nil {
		return
	}
	_, err = mongo.Write(Database, QuestionCollection, *requestDto)
	return
}

func (db DB) SeedQuestions(questions []Question) (err error) {
	var questionsDto []interface{}
	for _, t := range questions {
		questionsDto = append(questionsDto, t)
	}
	_, err = mongo.WriteMany(Database, QuestionCollection, questionsDto)
	return err
}

func (db DB) CreateAnswer(answer Answer) (err error) {
	requestDto, err := mongo.Document(answer)
	if err != nil {
		return
	}
	_, err = mongo.Write(Database, AnswerCollection, *requestDto)
	return
}

func (db DB) SeedAnswers(answers []Answer) (err error) {
	var answersDto []interface{}
	for _, answer := range answers {
		answersDto = append(answersDto, answer)
	}
	_, err = mongo.WriteMany(Database, AnswerCollection, answersDto)
	return err
}

func (db DB) CreateTeams(teams []Team) (err error) {
	var teamsDto []interface{}
	for _, t := range teams {
		teamsDto = append(teamsDto, t)
	}
	_, err = mongo.WriteMany(Database, TeamCollection, teamsDto)
	return
}

func (db DB) CreateSnapshot(snapshot Snapshot) error {
	return db.Create(snapshot, SnapshotCollection)
}

func (db DB) CreatePlayer(player Player) error {
	return db.Create(player, PlayerCollection)
}

func (db DB) CreateMatch(match Game) error {
	return db.Create(match, MatchCollection)
}

func (db DB) CreateSubscriber(subscriber Subscriber) error {
	return db.Create(subscriber, SubscriberCollection)
}

func (db DB) UpdateMatchQuestions(match Game, question Question) error {
	match.Tags = append(match.Tags, question.Tag)
	match.Active = true
	return db.UpdateMatch(match)
}

func (db DB) UpdateMatch(match Game) (err error) {
	requestDto, err := mongo.Document(match)
	if err != nil {
		return
	}
	_, err = mongo.Update(Database, MatchCollection, bson.M{"_id": match.Id}, bson.D{primitive.E{Key: "$set", Value: *requestDto}})
	return
}

func (db DB) UpdateTeam(team Team) (err error) {
	requestDto, err := mongo.Document(team)
	if err != nil {
		return
	}
	_, err = mongo.Update(Database, TeamCollection, bson.M{"_id": team.Id}, bson.D{primitive.E{Key: "$set", Value: *requestDto}})
	return
}

func (db DB) DeleteSubscriber(playerId string) (err error) {
	_, err = mongo.Delete(Database, SubscriberCollection, bson.M{
		"player_id": playerId,
		"active":    true})
	return
}
