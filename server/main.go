package main

import (
	"log"
	"net/http"
	"os"

	"github.com/atulanand206/find/server/core"
	"github.com/atulanand206/go-mongo"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	// Register the MongoDB cloud atlas database.
	mongoClientId := os.Getenv("MONGO_CLIENT_ID")
	mongo.ConfigureMongoClient(mongoClientId)

	// Register the endpoints exposed from the service.
	routes := core.Routes()

	// Serve the routes.
	handler := http.HandlerFunc(routes.ServeHTTP)
	port := os.Getenv("PORT")
	cert := os.Getenv("SSL_CERT")
	key := os.Getenv("SSL_KEY")
	if cert == "" || key == "" {
		log.Fatal("please add ssl certificates for successful connection.")
	}
	err := http.ListenAndServeTLS(":"+port, cert, key, handler)
	if err != nil {
		log.Fatal(err)
	}
}
