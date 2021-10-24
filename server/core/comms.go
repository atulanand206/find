package core

type Action int

const (
	BEGIN Action = iota
	SPECS
	JOIN
	WATCH
	CREATE
	START
	HINT
	RIGHT
	PASS
	NEXT
	REVEAL
	SCORE
	REFRESH
	S_REFRESH
	Over
	Extend
	ACTIVE
	S_ACTIVE
	S_GAME
	S_PLAYER
	S_START
	S_HINT
	S_RIGHT
	S_PASS
	S_NEXT
	S_SCORE
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
		"CREATE",
		"START",
		"HINT",
		"RIGHT",
		"PASS",
		"NEXT",
		"REVEAL",
		"SCORE",
		"REFRESH",
		"S_REFRESH",
		"over",
		"extend",
		"ACTIVE",
		"S_ACTIVE",
		"S_GAME",
		"S_PLAYER",
		"S_START",
		"S_HINT",
		"S_RIGHT",
		"S_PASS",
		"S_NEXT",
		"S_SCORE",
		"s_question",
		"s_answer",
		"s_over",
		"failure",
	}[action]
}

type Role int

const (
	QUIZMASTER Role = iota
	PLAYER
	AUDIENCE
	TEAM
)

func (role Role) String() string {
	return []string{
		"QUIZMASTER",
		"PLAYER",
		"AUDIENCE",
		"TEAM",
	}[role]
}
