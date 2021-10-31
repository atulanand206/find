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
	// Instance variable to store the DB Snapshot Collection name.
	SnapshotCollection string
	// Instance variable to store the DB Team Collection name.
	TeamCollection string
	// Instance variable to store the DB Player Collection name.
	PlayerCollection     string
	SubscriberCollection string

	CommsHub   *Hub
	Controller Service
	Targe      Target

	Db              DB
	InstanceCreator Creator
	MessageCreator  WebsocketMessageCreator
	ErrorCreator    ErrorMessageCreator
)

// Add handlers and interceptors to the endpoints.
func Routes() *http.ServeMux {
	Database = os.Getenv("GAME_DATABASE")
	MatchCollection = "matches"
	QuestionCollection = "questions"
	AnswerCollection = "answers"
	SnapshotCollection = "snapshots"
	TeamCollection = "teams"
	PlayerCollection = "players"
	IndexCollection = "indexes"
	SubscriberCollection = "subscribers"

	// Interceptor chain for attaching to the requests.
	chain := net.MiddlewareChain{
		net.ApplicationJsonInterceptor(),
		net.AuthenticationInterceptor(),
	}

	// Interceptor chain with only PUT method.
	getChain := chain.Add(net.CorsInterceptor(http.MethodGet))
	putChain := chain.Add(net.CorsInterceptor(http.MethodPut))
	// Interceptor chain with only POST method.
	postChain := chain.Add(net.CorsInterceptor(http.MethodPost))

	Db = DB{}
	Targe = Target{}
	validator := Validator{}

	authSvc := AuthService{}
	matchSvc := MatchService{crud: MatchCrud{db: Db}}
	playerSvc := PlayerService{crud: PlayerCrud{db: Db}}
	subscriberService := SubscriberService{crud: SubscriberCrud{db: Db}, target: Targe}
	teamSvc := TeamService{crud: TeamCrud{}, subscriberService: subscriberService}
	snapshotService := SnapshotService{crud: SnapshotCrud{db: Db}}
	questionService := QuestionService{crud: QuestionCrud{db: Db}}

	PermissionHandler := PermissionHandler{crud: PermissionCrud{}}
	MatchHandler := MatchHandler{matchService: matchSvc}

	Controller = Service{
		authService:       authSvc,
		matchService:      matchSvc,
		teamService:       teamSvc,
		playerService:     playerSvc,
		subscriberService: subscriberService,
		questionService:   questionService,
		snapshotService:   snapshotService,
		validator:         validator,
	}

	InstanceCreator = Creator{}
	MessageCreator = WebsocketMessageCreator{}
	ErrorCreator = ErrorMessageCreator{}

	router := http.NewServeMux()

	router.HandleFunc("/quizzes/active", getChain.Handler(MatchHandler.HandlerActiveQuizzes))

	router.HandleFunc("/permission/create", postChain.Handler(PermissionHandler.HandlerCreatePermission))
	router.HandleFunc("/permissions", getChain.Handler(PermissionHandler.HandlerFindPermissions))

	router.HandleFunc("/test", postChain.Handler(HandlerTestAPI))
	router.HandleFunc("/question/add", postChain.Handler(HandlerAddQuestion))
	router.HandleFunc("/questions/seed", putChain.Handler(HandlerSeedQuestions))

	// Register the websocket connection hub.
	CommsHub = NewHub()
	go CommsHub.Run()
	router.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		ServeWebsocket(CommsHub, w, r)
	})
	return router
}
