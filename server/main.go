package main

import (
	"net/http"
	"os"

	"github.com/atulanand206/find/server/routes"
	"github.com/atulanand206/go-mongo"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	// Register the MongoDB cloud atlas database.
	mongoClientId := os.Getenv("MONGO_CLIENT_ID")
	mongo.ConfigureMongoClient(mongoClientId)

	// Register the endpoints exposed from the service.
	routes := routes.Routes()
	handler := http.HandlerFunc(routes.ServeHTTP)
	http.ListenAndServe(":5000", handler)
}