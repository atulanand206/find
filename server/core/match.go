package core

const (
	PlayersInATeam    = 4
	TeamsInAMatch     = 4
	QuestionsInAMatch = 20
)

type (
	Game struct {
		Id         string   `json:"id" bson:"_id"`
		QuizMaster Player   `json:"quizmaster" bson:"quizmaster"`
		Tags       []string `bson:"tags"`
		Specs      Specs    `json:"specs" bson:"specs"`
		Active     bool     `json:"active" bson:"active"`
	}

	Specs struct {
		Teams     int `json:"teams"`
		Players   int `json:"players"`
		Questions int `json:"questions"`
		Rounds    int `json:"rounds"`
		Points    int `json:"points"`
	}

	Team struct {
		Id     string `json:"id" bson:"_id"`
		QuizId string `json:"quiz_id" bson:"quiz_id"`
		Name   string `json:"name" bson:"name"`
		Score  int    `json:"score" bson:"score"`
	}

	TeamRoster struct {
		Id      string   `json:"id"`
		Name    string   `json:"name"`
		Players []Player `json:"players"`
		Score   int      `json:"score" bson:"score"`
	}

	Player struct {
		Id    string `json:"id,omitempty" bson:"_id"`
		Name  string `json:"name,omitempty" bson:"name"`
		Email string `json:"email,omitempty" bson:"email"`
	}

	TeamPlayerRequest struct {
		TeamId   string `json:"team_id" bson:"team_id"`
		PlayerId string `json:"player_id" bson:"player_id"`
	}

	TeamPlayer struct {
		Id       string `json:"id" bson:"_id"`
		TeamId   string `json:"team_id" bson:"team_id"`
		PlayerId string `json:"player_id" bson:"player_id"`
	}

	Index struct {
		Id  string `json:"id" bson:"_id"`
		Tag string `json:"tag" bson:"tag"`
	}

	IndexWrapper struct {
		Indexes []Index `json:"indexes" bson:"indexes"`
	}

	IndexStore struct {
		Indexes []string `json:"indexes" bson:"indexes"`
	}

	QuestionBank struct {
		Questions []NewQuestion `json:"questions" bson:"questions"`
	}

	NewQuestion struct {
		Statements []string `json:"statements"`
		Answer     []string `json:"answer"`
	}

	Question struct {
		Id         string   `json:"id" bson:"_id"`
		Statements []string `json:"statements" bson:"statements"`
		Tag        string   `json:"-" bson:"tag"`
	}

	Answer struct {
		Id         string   `json:"id" bson:"_id"`
		QuestionId string   `json:"question_id" bson:"question_id"`
		Answer     []string `json:"answer" bson:"answer"`
		Hint       []string `json:"hints" bson:"hints"`
	}

	Snapshot struct {
		QuizId     string   `json:"quiz_id" bson:"quiz_id"`
		RoundNo    int      `json:"round_no" bson:"round_no"`
		QuestionNo int      `json:"question_no" bson:"question_no"`
		QuestionId string   `json:"question_id" bson:"question_id"`
		TeamSTurn  string   `json:"team_s_turn" bson:"team_s_turn"`
		EventType  string   `json:"event_type" bson:"event_type"`
		Score      int      `json:"score" bson:"score"`
		Timestamp  string   `json:"timestamp" bson:"timestamp"`
		Content    []string `json:"content" bson:"content"`
	}

	AddQuestionRequest struct {
		Question NewQuestion `json:"question"`
		Tag      string      `json:"tag"`
	}

	AddQuestionResponse struct {
		QuestionId string `json:"question_id"`
		AnswerId   string `json:"answer_id"`
	}

	EnterGameRequest struct {
		Person Player `json:"person"`
		QuizId string `json:"quiz_id"`
		TeamId string `json:"team_id"`
	}

	EnterGameResponse struct {
		Quiz         Game         `json:"quiz"`
		Roster       []TeamRoster `json:"teams"`
		Snapshot     Snapshot     `json:"snapshot"`
		PlayerTeamId string       `json:"player_team_id"`
	}

	CreateGameRequest struct {
		Quizmaster Player `json:"quizmaster"`
		Specs      Specs  `json:"specs"`
	}

	StartGameRequest struct {
		QuizId string `json:"quiz_id"`
	}

	StartGameResponse struct {
		QuizId   string       `json:"quiz_id"`
		Roster   []TeamRoster `json:"teams"`
		Question Question     `json:"question"`
		Snapshot Snapshot     `json:"snapshot"`
	}

	GameSnapRequest struct {
		QuizId     string `json:"quiz_id"`
		TeamSTurn  string `json:"team_s_turn"`
		QuestionId string `json:"question_id"`
	}

	NextQuestionRequest struct {
		QuizId         string `json:"quiz_id"`
		TeamSTurn      string `json:"team_s_turn"`
		LastQuestionId string `json:"question_id"`
	}

	ScoreRequest struct {
		QuizId string `json:"quiz_id"`
	}

	ScoreResponse struct {
		QuizId    string     `json:"quiz_id"`
		Snapshots []Snapshot `json:"snapshots"`
	}

	Subscriber struct {
		Tag      string `json:"tag" bson:"tag"`
		PlayerId string `json:"playerId" bson:"playerId"`
		Role     string `json:"role" bson:"role"`
		Active   bool   `json:"active" bson:"active"`
	}
)
