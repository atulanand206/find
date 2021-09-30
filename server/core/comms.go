package core

type Action int

const (
	BEGIN Action = iota
	SPECS
	JOIN
	WATCH
	Start
	REVEAL
	Attempt
	Score
	Next
	Over
	Extend
	S_GAME
	S_PLAYER
	S_START
	S_Question
	S_Answer
	S_Over
	Failure
)

func (action Action) String() string {
	return []string{
		"BEGIN",
		"SPECS",
		"JOIN",
		"WATCH",
		"start",
		"REVEAL",
		"attempt",
		"score",
		"next",
		"over",
		"extend",
		"S_GAME",
		"S_PLAYER",
		"S_START",
		"s_question",
		"s_answer",
		"s_over",
		"failure",
	}[action]
}
