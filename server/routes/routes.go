package routes

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/atulanand206/go-mongo"
	net "github.com/atulanand206/go-network"
	"go.mongodb.org/mongo-driver/bson"
	mg "go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/x/mongo/driver/uuid"
)

const (
	Err_RequestNotDecoded      = "Request can't be decoded."
	Err_MatchNotPresent        = "Match with the given info does not exist."
	Err_MatchNotDecoded        = "Match information can't be decoded."
	Err_MatchRequestNotCreated = "New match request can't be created."
	Err_MatchNotCreated        = "New match can't be created."
	Err_MatchNotUpdated        = "Match can't be updated."
	Err_PlayerNotCreated       = "Player can't be created."
	Err_PlayerNotDecoded       = "Player information can't be decoded."
	Err_PlayerNotPresent       = "Player with the given info does not exist."
)

type (
	Player struct {
		Id    uuid.UUID `json:"id" bson:"_id"`
		Name  string    `json:"name" bson:"name"`
		Email string    `json:"email" bson:"email"`
	}

	Game struct {
		Id         uuid.UUID `json:"id" bson:"_id"`
		Players    []Player  `json:"players" bson:"players"`
		QuizMaster Player    `json:"quizmaster" bson:"quizmaster"`
	}

	EnterGameRequest struct {
		Person  Player `json:"person"`
		MatchId string `json:"match_id"`
	}
)

var (
	// Instance variable to store the Database name.
	Database string
	// Instance variable to store the DB Match Collection name.
	MatchCollection string
	// Instance variable to store the DB Question Collection name.
	QuestionCollection string
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
	router.HandleFunc("/question/next", getChain.Handler(HandlerNextQuestion))
	router.HandleFunc("/question/verify", getChain.Handler(HandlerVerifyAnswer))
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
	player.Id, _ = uuid.New()
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
		match.Id, _ = uuid.New()
		match.QuizMaster = person
		requestDto, _ := mongo.Document(match)
		insertMatchResult, er := mongo.Write(Database, MatchCollection, *requestDto)
		if er != nil {
			http.Error(w, Err_MatchNotCreated, http.StatusInternalServerError)
			return
		}
		matchId = fmt.Sprint(insertMatchResult.InsertedID)
	} else {
		match.Players = append(match.Players, person)
		requestDto, _ := mongo.Document(match)
		_, err = mongo.Update(Database, MatchCollection, bson.M{"_id": matchId}, *requestDto)
		if err != nil {
			http.Error(w, Err_MatchNotUpdated, http.StatusInternalServerError)
			return
		}
	}
	dto = mongo.FindOne(Database, MatchCollection, bson.M{"_id": matchId})
	if err != nil {
		http.Error(w, Err_MatchNotDecoded, http.StatusInternalServerError)
		return
	}
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

}

func HandlerNextQuestion(w http.ResponseWriter, r *http.Request) {
}

func HandlerVerifyAnswer(w http.ResponseWriter, r *http.Request) {
}
