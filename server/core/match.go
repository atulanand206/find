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
		TeamSTurn  string     `json:"team_s_turn" bson:"team_s_turn"`
		Ready      bool       `json:"ready" bson:"ready"`
	}

	Specs struct {
		Teams     int `json:"teams"`
		Players   int `json:"players"`
		Questions int `json:"questions"`
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
		Answer     string   `json:"answer"`
	}

	Question struct {
		Id         string   `json:"id" bson:"_id"`
		Statements []string `json:"statements" bson:"statements"`
		Tag        string   `json:"-" bson:"tag"`
	}

	Answer struct {
		Id         string `json:"id" bson:"_id"`
		QuestionId string `json:"question_id" bson:"question_id"`
		Answer     string `json:"answer" bson:"answer"`
		Hint       string `json:"hint" bson:"hint"
		`
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
	}

	GameSnapRequest struct {
		QuizId     string `json:"quiz_id"`
		TeamSTurn  string `json:"team_s_turn"`
		QuestionId string `json:"question_id"`
	}

	HintRevealResponse struct {
		QuizId     string `json:"quiz_id"`
		TeamSTurn  string `json:"team_s_turn"`
		QuestionId string `json:"question_id"`
		Hint       string `json:"hint"`
	}

	AnswerRevealResponse struct {
		QuizId     string `json:"quiz_id"`
		TeamSTurn  string `json:"team_s_turn"`
		QuestionId string `json:"question_id"`
		Answer     string `json:"answer"`
	}

	GamePassResponse struct {
		QuizId     string `json:"quiz_id"`
		TeamSTurn  string `json:"team_s_turn"`
		QuestionId string `json:"question_id"`
	}

	GameNextResponse struct {
		QuizId         string   `json:"quiz_id"`
		TeamSTurn      string   `json:"team_s_turn"`
		LastQuestionId string   `json:"last_question_id"`
		Question       Question `json:"question"`
	}

	NextQuestionRequest struct {
		QuizId         string `json:"quiz_id"`
		TeamSTurn      string `json:"team_s_turn"`
		LastQuestionId string `json:"last_question_id"`
	}

	FindAnswerRequest struct {
		QuestionId string `json:"question_id"`
	}

	FindAnswerResponse struct {
		QuestionId string `json:"question_id"`
		Answer     string `json:"answer"`
	}
)
