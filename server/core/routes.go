package core

import (
	"net/http"
	"os"

	net "github.com/atulanand206/go-network"
)

var (
	// Instance variable to store the Database name.
	Database string
	// Instance variable to store the DB Match Collection name.
	MatchCollection string
	// Instance variable to store the DB Index Collection name.
	IndexCollection string
	// Instance variable to store the DB Question Collection name.
	QuestionCollection string
	// Instance variable to store the DB Answer Collection name.
	AnswerCollection string
	// Instance variable to store the DB Match Collection name.
	PlayerCollection string

	hub *Hub
)

// Add handlers and interceptors to the endpoints.
func Routes() *http.ServeMux {
	Database = os.Getenv("GAME_DATABASE")
	MatchCollection = os.Getenv("MATCH_COLLECTION")
	QuestionCollection = os.Getenv("QUESTION_COLLECTION")
	AnswerCollection = os.Getenv("ANSWER_COLLECTION")
	PlayerCollection = os.Getenv("PLAYER_COLLECTION")
	IndexCollection = os.Getenv("INDEX_COLLECTION")

	// Interceptor chain for attaching to the requests.
	chain := net.MiddlewareChain{
		net.ApplicationJsonInterceptor(),
		// net.AuthenticationInterceptor(),
	}

	// Interceptor chain with only PUT method.
	putChain := chain.Add(net.CorsInterceptor(http.MethodPut))
	// Interceptor chain with only POST method.
	postChain := chain.Add(net.CorsInterceptor(http.MethodPost))

	router := http.NewServeMux()

	router.HandleFunc("/match/begin", postChain.Handler(HandlerBeginGame))
	router.HandleFunc("/match/enter", postChain.Handler(HandlerEnterGame))
	router.HandleFunc("/match/start", postChain.Handler(HandlerStartGame))

	router.HandleFunc("/question/add", postChain.Handler(HandlerAddQuestion))
	router.HandleFunc("/question/next", postChain.Handler(HandlerNextQuestion))
	router.HandleFunc("/question/answer", postChain.Handler(HandlerFindAnswer))

	router.HandleFunc("/questions/seed", putChain.Handler(HandlerSeedQuestions))

	// Register the websocket connection hub.
	hub = NewHub()
	go hub.Run()
	router.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		ServeWebsocket(hub, w, r)
	})
	return router
}
