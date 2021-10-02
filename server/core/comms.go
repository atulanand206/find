package core

type Action int

const (
	BEGIN Action = iota
	SPECS
	JOIN
	WATCH
	START
	HINT
	RIGHT
	PASS
	NEXT
	REVEAL
	Score
	Over
	Extend
	S_GAME
	S_PLAYER
	S_START
	S_HINT
	S_RIGHT
	S_PASS
	S_NEXT
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
		"START",
		"HINT",
		"RIGHT",
		"PASS",
		"NEXT",
		"REVEAL",
		"score",
		"over",
		"extend",
		"S_GAME",
		"S_PLAYER",
		"S_START",
		"S_HINT",
		"S_RIGHT",
		"S_PASS",
		"S_NEXT",
		"s_question",
		"s_answer",
		"s_over",
		"failure",
	}[action]
}
