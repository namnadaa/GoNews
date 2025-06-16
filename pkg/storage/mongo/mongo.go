package mongo

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MongoStorage - хранилище данных с подключением к БД MongoDB.
type MongoStorage struct {
	db             *mongo.Client
	databaseName   string
	collectionName string
}

// NewMongoStorage создает новое подключение к базе данных MongoDB.
func NewMongoStorage(content, dbName, colName string) (*MongoStorage, error) {
	mongoOpts := options.Client().ApplyURI(content)
	db, err := mongo.Connect(context.Background(), mongoOpts)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to MongoDB: %v", err)
	}

	err = db.Ping(context.Background(), nil)
	if err != nil {
		return nil, fmt.Errorf("cannot ping MongoDB: %v", err)
	}

	s := MongoStorage{
		db:             db,
		databaseName:   dbName,
		collectionName: colName,
	}
	return &s, nil
}

// CloseMongo закрывает соединение с базой данных MongoDB.
func (ms *MongoStorage) CloseMongo() error {
	err := ms.db.Disconnect(context.Background())
	if err != nil {
		return fmt.Errorf("failed to close MongoDB connection: %v", err)
	}
	return nil
}
