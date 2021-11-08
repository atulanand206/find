package comms

import (
	"net/http"

	"github.com/atulanand206/find/server/core/db"
	"github.com/atulanand206/find/server/core/services"
	"github.com/atulanand206/go-mongo"
	net "github.com/atulanand206/go-network"
)

var (
	Db         db.DB
	CommsHub   *Hub
	Controller services.Service
)

// Add handlers and interceptors to the endpoints.
func Routes(mongoClientId string, database string) *http.ServeMux {
	mongo.ConfigureMongoClient(mongoClientId)
	db.Database = database

	// Interceptor chain for attaching to the requests.
	chain := net.MiddlewareChain{
		net.ApplicationJsonInterceptor(),
		// net.AuthenticationInterceptor(),
	}

	// Interceptor chain with only PUT method.
	getChain := chain.Add(net.CorsInterceptor(http.MethodGet))
	putChain := chain.Add(net.CorsInterceptor(http.MethodPut))
	// Interceptor chain with only POST method.
	postChain := chain.Add(net.CorsInterceptor(http.MethodPost))

	Db = db.DB{}
	Controller = services.Init(Db)

	PermissionHandler := PermissionHandler{crud: db.PermissionCrud{}}
	MatchHandler := MatchHandler{matchService: Controller.MatchService}

	router := http.NewServeMux()

	router.HandleFunc("/quizzes/active", postChain.Handler(MatchHandler.HandlerActiveQuizzes))

	router.HandleFunc("/permission/create", postChain.Handler(PermissionHandler.HandlerCreatePermission))
	router.HandleFunc("/permissions", getChain.Handler(PermissionHandler.HandlerFindPermissions))

	router.HandleFunc("/test", postChain.Handler(HandlerTestAPI))
	router.HandleFunc("/question/add", postChain.Handler(HandlerAddQuestion))
	router.HandleFunc("/questions/seed", putChain.Handler(HandlerSeedQuestions))

	// Register the websocket connection hub.
	CommsHub = NewHub(Controller)
	go CommsHub.Run()
	router.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		ServeWebsocket(CommsHub, w, r)
	})
	return router
}
