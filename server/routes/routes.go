package routes

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"

	"github.com/atulanand206/go-mongo"
	net "github.com/atulanand206/go-network"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	mg "go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	PlayersCount = 2

	Err_QuestionsNotSeeded     = "Questions can't be seeded."
	Err_AnswersNotSeeded       = "Answers can't be seeded."
	Err_RequestNotDecoded      = "Request can't be decoded."
	Err_MatchNotPresent        = "Match with the given info does not exist."
	Err_MatchNotDecoded        = "Match information can't be decoded."
	Err_MatchRequestNotCreated = "New match request can't be created."
	Err_MatchNotCreated        = "New match can't be created."
	Err_MatchNotUpdated        = "Match can't be updated."
	Err_PlayerNotCreated       = "Player can't be created."
	Err_PlayerNotDecoded       = "Player information can't be decoded."
	Err_PlayerNotPresent       = "Player with the given info does not exist."
	Err_QuizmasterCantPlay     = "Quizmaster can't join the match as a player."
	Err_WaitingForPlayers      = "Waiting for more players to join."
	Err_QuestionsRequestFailed = "Questions count request failed."
	Err_QuestionNotCreated     = "Question can't be created."
	Err_AnswerNotCreated       = "Answer can't be created."
	Err_AnswerNotPresent       = "Answer does not exist."
)

type (
	Player struct {
		Id    string `json:"id" bson:"_id"`
		Name  string `json:"name" bson:"name"`
		Email string `json:"email" bson:"email"`
	}

	Game struct {
		Id         string   `json:"id" bson:"_id"`
		Players    []Player `json:"players" bson:"players"`
		QuizMaster Player   `json:"quizmaster" bson:"quizmaster"`
	}

	EnterGameRequest struct {
		Person  Player `json:"person"`
		MatchId string `json:"match_id"`
	}

	StartGameRequest struct {
		MatchId string `json:"match_id"`
	}

	StartGameResponse struct {
		Match  Game     `json:"game"`
		Prompt Question `json:"prompt"`
	}

	Question struct {
		Id         string   `json:"id" bson:"_id"`
		Statements []string `json:"statements" bson:"statements"`
		Tag        string   `bson:"tag"`
	}

	Answer struct {
		Id         string `json:"id" bson:"_id"`
		QuestionId string `json:"question_id" bson:"question_id"`
		Answer     string `json:"answer" bson:"answer"`
	}

	FindAnswerRequest struct {
		QuestionId string `json:"question_id"`
	}

	FindAnswerResponse struct {
		QuestionId string `json:"question_id"`
		Answer     string `json:"answer"`
	}

	WebsocketMessage struct {
		Person Player `json:"person"`
	}
)

var (
	// Instance variable to store the Database name.
	Database string
	// Instance variable to store the DB Match Collection name.
	MatchCollection string
	// Instance variable to store the DB Question Collection name.
	QuestionCollection string
	// Instance variable to store the DB Answer Collection name.
	AnswerCollection string
	// Instance variable to store the DB Match Collection name.
	PlayerCollection string
)

// Interceptor which checks if the header has the correct accessToken.
func LoggingInterceptor() net.MiddlewareInterceptor {
	return func(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
		fmt.Println(r)
		next(w, r)
	}
}

// Add handlers and interceptors to the endpoints.
func Routes() *http.ServeMux {
	Database = os.Getenv("GAME_DATABASE")
	MatchCollection = os.Getenv("MATCH_COLLECTION")
	QuestionCollection = os.Getenv("QUESTION_COLLECTION")
	AnswerCollection = os.Getenv("ANSWER_COLLECTION")
	PlayerCollection = os.Getenv("PLAYER_COLLECTION")

	// Interceptor chain for attaching to the requests.
	chain := net.MiddlewareChain{
		net.ApplicationJsonInterceptor(),
		// net.AuthenticationInterceptor(),
	}

	// Interceptor chain with only GET method.
	getChain := chain.Add(net.CorsInterceptor(http.MethodGet))
	// Interceptor chain with only POST method.
	postChain := chain.Add(net.CorsInterceptor(http.MethodPost))

	router := http.NewServeMux()
	router.HandleFunc("/enter", postChain.Handler(HandlerEnterGame))
	router.HandleFunc("/start", postChain.Handler(HandlerStart))
	router.HandleFunc("/question/add", postChain.Handler(HandlerAddQuestion))
	router.HandleFunc("/questions/seed", getChain.Handler(HandlerSeedQuestions))
	router.HandleFunc("/question/verify", getChain.Handler(HandlerFindAnswer))
	router.HandleFunc("/ws", chain.Handler(HandlerWebSockets))
	go handleMessages()
	return router
}

func HandlerEnterGame(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var requestBody EnterGameRequest
	err := decoder.Decode(&requestBody)
	if err != nil {
		http.Error(w, Err_RequestNotDecoded, http.StatusInternalServerError)
		return
	}
	player, err := CreateOrFindPlayer(w, requestBody.Person)
	if err != nil {
		return
	}
	match, err := CreateOrUpdateMatch(w, requestBody.MatchId, player)
	if err != nil {
		return
	}
	json.NewEncoder(w).Encode(match)
}

func CreateOrFindPlayer(w http.ResponseWriter, playerRequest Player) (player Player, err error) {
	dto := mongo.FindOne(Database, PlayerCollection, bson.M{"email": playerRequest.Email})
	err = dto.Err()
	if err == nil {
		player, err = DecodePlayer(dto)
		return
	}
	player = playerRequest
	player.Id = uuid.New().String()
	requestDto, _ := mongo.Document(player)
	if err != nil {
		_, err = mongo.Write(Database, PlayerCollection, *requestDto)
		if err != nil {
			http.Error(w, Err_PlayerNotCreated, http.StatusInternalServerError)
			return
		}
	}
	dto = mongo.FindOne(Database, PlayerCollection, bson.M{"email": player.Email})
	if err != nil {
		http.Error(w, Err_PlayerNotPresent, http.StatusInternalServerError)
		return
	}
	player, err = DecodePlayer(dto)
	return
}

// Decodes a mongo db single result into an user object.
func DecodePlayer(document *mg.SingleResult) (v Player, err error) {
	var player Player
	if err = document.Decode(&player); err != nil {
		return player, err
	}
	return player, err
}

func CreateOrUpdateMatch(w http.ResponseWriter, matchId string, person Player) (match Game, err error) {
	dto := mongo.FindOne(Database, MatchCollection, bson.M{"_id": matchId})
	err = dto.Err()
	if err == nil {
		match, err = DecodeMatch(dto)
	}
	for _, v := range match.Players {
		if v.Id == person.Id {
			return
		}
	}
	if err != nil {
		match.Id = uuid.New().String()
		match.Players = make([]Player, 0)
		match.QuizMaster = person
		requestDto, _ := mongo.Document(match)
		insertMatchResult, er := mongo.Write(Database, MatchCollection, *requestDto)
		err = nil
		if er != nil {
			http.Error(w, Err_MatchNotCreated, http.StatusInternalServerError)
			return
		}
		matchId = fmt.Sprint(insertMatchResult.InsertedID)
	} else {
		if match.QuizMaster.Id == person.Id {
			err = errors.New(Err_QuizmasterCantPlay)
			http.Error(w, Err_QuizmasterCantPlay, http.StatusInternalServerError)
			return
		}
		match.Players = append(match.Players, person)
		requestDto, _ := mongo.Document(match)
		_, err = mongo.Update(Database, MatchCollection, bson.M{"_id": matchId}, bson.D{primitive.E{Key: "$set", Value: *requestDto}})
		if err != nil {
			http.Error(w, Err_MatchNotUpdated, http.StatusInternalServerError)
			return
		}
	}
	dto = mongo.FindOne(Database, MatchCollection, bson.M{"_id": matchId})
	match, err = DecodeMatch(dto)
	return
}

// Decodes a mongo db single result into an user object.
func DecodeMatch(document *mg.SingleResult) (v Game, err error) {
	var game Game
	if err = document.Decode(&game); err != nil {
		return game, err
	}
	return game, err
}

func HandlerStart(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var requestBody StartGameRequest
	err := decoder.Decode(&requestBody)
	if err != nil {
		http.Error(w, Err_RequestNotDecoded, http.StatusInternalServerError)
		return
	}
	dto := mongo.FindOne(Database, MatchCollection, bson.M{"_id": requestBody.MatchId})
	err = dto.Err()
	if err != nil {
		http.Error(w, Err_MatchNotPresent, http.StatusInternalServerError)
		return
	}
	match, err := DecodeMatch(dto)
	if len(match.Players) != PlayersCount {
		http.Error(w, Err_WaitingForPlayers, http.StatusInternalServerError)
		return
	}
	question, err := FindQuestion(w)
	if err != nil {
		return
	}
	var response StartGameResponse
	response.Match = match
	response.Prompt = question
	json.NewEncoder(w).Encode(response)
}

func QuestionsCount() (int, error) {
	cursor, err := mongo.Find(Database, QuestionCollection, bson.M{}, &options.FindOptions{})
	if err != nil {
		return 0, err
	}
	elements, _ := cursor.Current.Elements()
	size := len(elements)
	return size, nil
}

func FindQuestion(w http.ResponseWriter) (question Question, err error) {
	return
}

func HandlerAddQuestion(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var requestBody NewQuestion
	err := decoder.Decode(&requestBody)
	if err != nil {
		http.Error(w, Err_RequestNotDecoded, http.StatusInternalServerError)
		return
	}
	var question Question
	question.Id = uuid.New().String()
	question.Statements = requestBody.Statements
	requestDto, _ := mongo.Document(question)
	insertedQuestionId, err := mongo.Write(Database, QuestionCollection, *requestDto)
	if err != nil {
		http.Error(w, Err_QuestionNotCreated, http.StatusInternalServerError)
		return
	}
	var answer Answer
	answer.Id = uuid.New().String()
	answer.QuestionId = fmt.Sprint(insertedQuestionId.InsertedID)
	answer.Answer = requestBody.Answer
	requestDto, _ = mongo.Document(answer)
	insertedAnswerId, err := mongo.Write(Database, AnswerCollection, *requestDto)
	if err != nil {
		http.Error(w, Err_AnswerNotCreated, http.StatusInternalServerError)
		return
	}
	var response AddQuestionResponse
	response.QuestionId = fmt.Sprint(insertedQuestionId.InsertedID)
	response.AnswerId = fmt.Sprint(insertedAnswerId.InsertedID)
	json.NewEncoder(w).Encode(response)
}

func HandlerFindAnswer(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var requestBody FindAnswerRequest
	err := decoder.Decode(&requestBody)
	if err != nil {
		http.Error(w, Err_RequestNotDecoded, http.StatusInternalServerError)
		return
	}
	dto := mongo.FindOne(Database, AnswerCollection, bson.M{"question_id": requestBody.QuestionId})
	err = dto.Err()
	if err != nil {
		http.Error(w, Err_AnswerNotPresent, http.StatusInternalServerError)
		return
	}
	answer, err := DecodeAnswer(dto)
	var response FindAnswerResponse
	response.QuestionId = requestBody.QuestionId
	response.Answer = answer.Answer
	json.NewEncoder(w).Encode(response)
}

// Decodes a mongo db single result into an user object.
func DecodeAnswer(document *mg.SingleResult) (v Answer, err error) {
	var answer Answer
	if err = document.Decode(&answer); err != nil {
		return answer, err
	}
	return answer, err
}
