package main

import (
	"GoNews/pkg/api"
	"GoNews/pkg/storage"
	"GoNews/pkg/storage/memdb"
	"GoNews/pkg/storage/mongo"
	"GoNews/pkg/storage/postgres"
	"log"
	"net/http"
	"os"
)

// Сервер GoNews.
type server struct {
	db  storage.Interface
	api *api.API
}

func main() {
	var srv server

	switch os.Getenv("STORAGE") {
	case "postgres":
		dsn := os.Getenv("POSTGRES_DSN")
		if dsn == "" {
			dsn = "postgres://myuser:mypassword@localhost:5434/mydb"

		}

		dbPostgres, err := postgres.NewPostgresStorage(dsn)
		if err != nil {
			log.Fatalf("failed to init postgres storage: %v", err)
		}

		srv.db = dbPostgres
		log.Println("Using PostgreSQL storage")
		defer dbPostgres.ClosePostgres()
	case "mongo":
		dsn := os.Getenv("MONGO_DSN")
		if dsn == "" {
			dsn = "mongodb://localhost:27017/"
		}

		dbName := os.Getenv("MONGO_DB")
		if dbName == "" {
			dbName = "data"
		}

		collectionName := os.Getenv("MONGO_COLLECTION")
		if collectionName == "" {
			collectionName = "languages"
		}

		dbMongo, err := mongo.NewMongoStorage(dsn, dbName, collectionName)
		if err != nil {
			log.Fatalf("failed to init mongo storage: %v", err)
		}

		srv.db = dbMongo
		log.Println("Using MongoDB storage")
		defer dbMongo.CloseMongo()
	case "memory", "":
		dbMemory := memdb.NewMemoryStorage()
		srv.db = dbMemory
		log.Println("Using in-memory storage")
	default:
		log.Fatalf("unknown storage backend: %s", os.Getenv("STORAGE"))
	}

	srv.api = api.New(srv.db)

	log.Fatal(http.ListenAndServe(":8080", srv.api.Router()))
}
