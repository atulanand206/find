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

	CommsHub        *Hub
	Controller      Service
	Targe           Target
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
		// net.AuthenticationInterceptor(),
	}

	// Interceptor chain with only PUT method.
	putChain := chain.Add(net.CorsInterceptor(http.MethodPut))
	// Interceptor chain with only POST method.
	postChain := chain.Add(net.CorsInterceptor(http.MethodPost))

	Db = DB{}
	questionService := QuestionService{db: Db}
	Targe = Target{}
	subscriberService := SubscriberService{db: Db, target: Targe}
	validator := Validator{}
	Controller = Service{
		matchService:      MatchService{db: Db},
		teamService:       TeamService{db: Db, subscriberService: subscriberService},
		playerService:     PlayerService{db: Db},
		subscriberService: subscriberService,
		questionService:   questionService,
		validator:         validator,
	}

	InstanceCreator = Creator{}
	MessageCreator = WebsocketMessageCreator{}
	ErrorCreator = ErrorMessageCreator{}

	router := http.NewServeMux()

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
