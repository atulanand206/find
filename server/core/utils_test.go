package core_test

import (
	"testing"

	"github.com/atulanand206/find/server/core"
	"github.com/atulanand206/go-mongo"
)

func Setup(t *testing.T) func(t *testing.T) {
	mongo.ConfigureMongoClient("mongodb://localhost:27017")
	core.Database = "binquiz-test"
	core.MatchCollection = "matches"
	core.QuestionCollection = "questions"
	core.AnswerCollection = "answers"
	core.SnapshotCollection = "snapshots"
	core.TeamCollection = "teams"
	core.PlayerCollection = "players"
	core.IndexCollection = "indexes"
	core.SubscriberCollection = "subscribers"
	return func(t *testing.T) {
	}
}
