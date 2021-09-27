package core

const (
	PlayersInATeam    = 4
	TeamsInAMatch     = 4
	QuestionsInAMatch = 20
)

type (
	Game struct {
		Id         string   `json:"id" bson:"_id"`
		Teams      []Team   `json:"teams" bson:"teams"`
		QuizMaster Player   `json:"quizmaster" bson:"quizmaster"`
		Tags       []string `bson:"tags"`
	}

	Team struct {
		Id      string   `json:"id" bson:"_id"`
		Name    string   `json:"name" bson:"name"`
		Players []Player `json:"players" bson:"players"`
	}

	Player struct {
		Id    string `json:"id" bson:"_id"`
		Name  string `json:"name" bson:"name"`
		Email string `json:"email" bson:"email"`
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
		Person  Player `json:"person"`
		TeamId  string `json:"team_id"`
		MatchId string `json:"match_id"`
	}

	StartGameRequest struct {
		MatchId string `json:"match_id"`
	}

	StartGameResponse struct {
		Match  Game       `json:"game"`
		Prompt []Question `json:"prompt"`
	}

	NextQuestionRequest struct {
		MatchId string `json:"match_id"`
		Limit   int    `json:"limit"`
		Types   int    `json:"types"`
	}

	FindAnswerRequest struct {
		QuestionId string `json:"question_id"`
	}

	FindAnswerResponse struct {
		QuestionId string `json:"question_id"`
		Answer     string `json:"answer"`
	}
)
