package main

import (
	"log"
	"net/http"
	"os"

	"github.com/atulanand206/find/server/core"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	// Register the MongoDB cloud atlas database.
	mongoClientId := os.Getenv("MONGO_CLIENT_ID")
	Database := os.Getenv("GAME_DATABASE")
	// Register the endpoints exposed from the service.
	routes := core.Routes(mongoClientId, Database)

	// Serve the routes.
	handler := http.HandlerFunc(routes.ServeHTTP)
	port := os.Getenv("PORT")
	cert := os.Getenv("SSL_CERT")
	key := os.Getenv("SSL_KEY")
	if cert == "" || key == "" {
		err := http.ListenAndServe(":"+port, handler)
		if err != nil {
			log.Fatal(err)
		}
	}
	err := http.ListenAndServeTLS(":"+port, cert, key, handler)
	if err != nil {
		log.Fatal(err)
	}
}
