package models

const (
	PlayersInATeam    = 4
	TeamsInAMatch     = 4
	QuestionsInAMatch = 20
)

type (
	Permission struct {
		PlayerId string `json:"player_id" bson:"player_id"`
	}

	Game struct {
		Id            string   `json:"id" bson:"_id"`
		QuizMaster    Player   `json:"quizmaster" bson:"quizmaster"`
		Tags          []string `bson:"tags"`
		Specs         Specs    `json:"specs" bson:"specs"`
		Active        bool     `json:"active" bson:"active"`
		Started       bool     `json:"started" bson:"started"`
		CanJoin       bool     `json:"can_join"`
		PlayersJoined int      `json:"players_joined"`
	}

	Specs struct {
		Name      string `json:"name" bson:"name"`
		Teams     int    `json:"teams" bson:"teams"`
		Players   int    `json:"players" bson:"players"`
		Questions int    `json:"questions" bson:"questions"`
		Rounds    int    `json:"rounds" bson:"rounds"`
		Points    int    `json:"points" bson:"points"`
	}

	Team struct {
		Id     string `json:"id" bson:"_id"`
		QuizId string `json:"quiz_id" bson:"quiz_id"`
		Name   string `json:"name" bson:"name"`
		Score  int    `json:"score" bson:"score"`
	}

	Player struct {
		Id    string `json:"id" bson:"_id"`
		Name  string `json:"name" bson:"name"`
		Email string `json:"email" bson:"email"`
	}

	Snapshot struct {
		QuizId     string       `json:"quiz_id" bson:"quiz_id"`
		RoundNo    int          `json:"round_no" bson:"round_no"`
		Roster     []TeamRoster `json:"teams" bson:"teams"`
		TeamSTurn  string       `json:"team_s_turn" bson:"team_s_turn"`
		QuestionNo int          `json:"question_no" bson:"question_no"`
		QuestionId string       `json:"question_id" bson:"question_id"`
		EventType  string       `json:"event_type" bson:"event_type"`
		Score      int          `json:"score" bson:"score"`
		Timestamp  string       `json:"timestamp" bson:"timestamp"`
		Question   []string     `json:"question" bson:"question"`
		Answer     []string     `json:"answer" bson:"answer"`
		Hint       []string     `json:"hint" bson:"hint"`
		CanPass    bool         `json:"can_pass"`
	}

	Subscriber struct {
		Tag      string `json:"tag" bson:"tag"`
		PlayerId string `json:"player_id" bson:"player_id"`
		Role     string `json:"role" bson:"role"`
		Active   bool   `json:"active" bson:"active"`
	}

	Index struct {
		Id  string `json:"id" bson:"_id"`
		Tag string `json:"tag" bson:"tag"`
	}

	IndexStore struct {
		Indexes []string `json:"indexes" bson:"indexes"`
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

	QuestionBank struct {
		Questions []NewQuestion `json:"questions"`
	}

	NewQuestion struct {
		Statements []string `json:"statements"`
		Answer     []string `json:"answer"`
	}

	AddQuestionRequest struct {
		Questions []NewQuestion `json:"questions"`
		Tag       string        `json:"tag"`
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

	TeamRoster struct {
		Id      string   `json:"id"`
		Name    string   `json:"name"`
		Players []Player `json:"players"`
		Score   int      `json:"score" bson:"score"`
	}

	CreateGameRequest struct {
		Quizmaster Player `json:"quizmaster"`
		Specs      Specs  `json:"specs"`
	}

	StartGameRequest struct {
		QuizId string `json:"quiz_id"`
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

	Request struct {
		Action     string `json:"action"`
		Specs      Specs  `json:"specs,omitempty"`
		Person     Player `json:"person,omitempty"`
		QuizId     string `json:"quiz_id,omitempty"`
		TeamId     string `json:"team_id,omitempty"`
		QuestionId string `json:"question_id,omitempty"`
	}

	GameResponse struct {
		Quiz     Game     `json:"quiz"`
		Snapshot Snapshot `json:"snapshot"`
		Role     string   `json:"role"`
	}
)

type SnapshotResponse struct {
	Action   string   `json:"action"`
	Snapshot Snapshot `json:"snapshot"`
}

type AuthenticationResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type LoginResponse struct {
	Player     Player                 `json:"player"`
	Tokens     AuthenticationResponse `json:"tokens"`
	Quizmaster bool                   `json:"quizmaster"`
}

type (
	WebsocketMessage struct {
		Action  string `json:"action"`
		Content string `json:"content"`
	}
)
