package core

import (
	"context"

	mg "go.mongodb.org/mongo-driver/mongo"
)

func DecodeIndex(document *mg.SingleResult) (v IndexWrapper, err error) {
	var index IndexWrapper
	if err = document.Decode(&index); err != nil {
		return index, err
	}
	return index, err
}

func DecodeMatch(document *mg.SingleResult) (v Game, err error) {
	var game Game
	if err = document.Decode(&game); err != nil {
		return game, err
	}
	return game, err
}

func DecodePlayer(document *mg.SingleResult) (v Player, err error) {
	var player Player
	if err = document.Decode(&player); err != nil {
		return player, err
	}
	return player, err
}

func DecodeIndexes(cursor *mg.Cursor) (indexes []Index, err error) {
	for cursor.Next(context.Background()) {
		var index Index
		err = cursor.Decode(&index)
		if err != nil {
			return
		}
		indexes = append(indexes, index)
	}
	return
}

func DecodeAnswer(document *mg.SingleResult) (v Answer, err error) {
	var answer Answer
	if err = document.Decode(&answer); err != nil {
		return answer, err
	}
	return answer, err
}
