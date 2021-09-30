package core

import (
	"github.com/atulanand206/go-mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func QuestionsCollections() []string {
	return []string{IndexCollection, QuestionCollection, AnswerCollection}
}

func CreateQuestionsCollections() (err error) {
	err = mongo.CreateCollections(Database, QuestionsCollections())
	return
}

func DropQuestionsCollections() (err error) {
	err = mongo.DropCollections(Database, QuestionsCollections())
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

func UpdateMatchPlayer(match Game, player Player, teamId string) (err error) {
	for i := 0; i < len(match.Teams); i++ {
		team := match.Teams[i]
		if team.Id == teamId {
			team.Players = append(team.Players, player)
			match.Teams[i] = team
		}
	}
	err = UpdateMatch(match)
	return
}

func UpdateMatchQuestions(match Game, questions []Question) (err error) {
	for _, question := range questions {
		match.Tags = append(match.Tags, question.Tag)
	}
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
