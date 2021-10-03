package core

const (
	PlayersInATeam    = 4
	TeamsInAMatch     = 4
	QuestionsInAMatch = 20
)

type (
	Game struct {
		Id         string     `json:"id" bson:"_id"`
		Teams      []TeamMini `json:"teams" bson:"teams"`
		QuizMaster Player     `json:"quizmaster" bson:"quizmaster"`
		Tags       []string   `bson:"tags"`
		Specs      Specs      `json:"specs" bson:"specs"`
	}

	Specs struct {
		Teams     int `json:"teams"`
		Players   int `json:"players"`
		Questions int `json:"questions"`
		Rounds    int `json:"rounds"`
		Points    int `json:"points"`
	}

	TeamMini struct {
		Id    string `json:"id" bson:"_id"`
		Name  string `json:"name" bson:"name"`
		Score int    `json:"score" bson:"score"`
	}

	Team struct {
		Id      string   `json:"id" bson:"_id"`
		Name    string   `json:"name" bson:"name"`
		Players []Player `json:"players" bson:"players"`
	}

	Player struct {
		Id    string `json:"id,omitempty" bson:"_id"`
		Name  string `json:"name,omitempty" bson:"name"`
		Email string `json:"email,omitempty" bson:"email"`
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
		Quiz  Game   `json:"quiz"`
		Teams []Team `json:"teams"`
	}

	CreateGameRequest struct {
		Quizmaster Player `json:"quizmaster"`
		Specs      Specs  `json:"specs"`
	}

	StartGameRequest struct {
		QuizId string `json:"quiz_id"`
	}

	StartGameResponse struct {
		QuizId   string   `json:"quiz_id"`
		Teams    []Team   `json:"teams"`
		Question Question `json:"question"`
		Snapshot Snapshot `json:"snapshot"`
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
)
