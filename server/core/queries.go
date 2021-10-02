package core

import (
	"github.com/atulanand206/go-mongo"
	"github.com/xorcare/pointer"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func FindMatch(matchId string) (match Game, err error) {
	dto := mongo.FindOne(Database, MatchCollection, bson.M{"_id": matchId})
	if err = dto.Err(); err != nil {
		return
	}
	match, err = DecodeMatch(dto)
	return
}

func FindQuestion(questionId string) (question Question, err error) {
	dto := mongo.FindOne(Database, MatchCollection, bson.M{"_id": questionId})
	if err = dto.Err(); err != nil {
		return
	}
	question, err = DecodeQuestion(dto)
	return
}

func FindPlayer(emailId string) (player Player, err error) {
	dto := mongo.FindOne(Database, PlayerCollection, bson.M{"email": emailId})
	if err = dto.Err(); err != nil {
		return
	}
	player, err = DecodePlayer(dto)
	return
}

func FindIndexForTag(tag string) (index Index, err error) {
	dto := mongo.FindOne(Database, IndexCollection, bson.M{"tag": tag})
	if err = dto.Err(); err != nil {
		return
	}
	index, err = DecodeIndex(dto)
	return
}

func FindAnswer(questionId string) (answer Answer, err error) {
	dto := mongo.FindOne(Database, AnswerCollection, bson.M{"question_id": questionId})
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
	teamIds := make([]string, 0)
	for _, v := range match.Teams {
		teamIds = append(teamIds, v.Id)
	}
	cursor, err := mongo.Find(Database, TeamCollection,
		bson.M{"_id": bson.M{"$in": teamIds}}, &options.FindOptions{})
	if err != nil {
		return
	}
	teams, err = DecodeTeams(cursor)
	return
}
