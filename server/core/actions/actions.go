package actions

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
	DROP
	SCORE
	REFRESH
	S_REFRESH
	FINISH
	Extend
	ACTIVE
	S_ACTIVE
	S_GAME
	S_JOIN
	S_PLAYER
	S_START
	S_HINT
	S_RIGHT
	S_PASS
	S_NEXT
	S_SCORE
	S_Question
	S_Answer
	S_OVER
	Failure
	EMPTY
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
		"DROP",
		"SCORE",
		"REFRESH",
		"S_REFRESH",
		"FINISH",
		"extend",
		"ACTIVE",
		"S_ACTIVE",
		"S_GAME",
		"S_JOIN",
		"S_PLAYER",
		"S_START",
		"S_HINT",
		"S_RIGHT",
		"S_PASS",
		"S_NEXT",
		"S_SCORE",
		"s_question",
		"s_answer",
		"S_OVER",
		"failure",
		"EMPTY",
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
