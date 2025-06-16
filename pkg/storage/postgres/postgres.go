package postgres

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"
)

// PostgresStorage - хранилище данных с подключением к БД PostgreSQL через пул соединений.
type PostgresStorage struct {
	db *pgxpool.Pool
}

// NewPostgresStorage создает новое подключение к базе данных PostgreSQL.
func NewPostgresStorage(content string) (*PostgresStorage, error) {
	db, err := pgxpool.Connect(context.Background(), content)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to PostgreSQL: %v", err)
	}

	err = db.Ping(context.Background())
	if err != nil {
		return nil, fmt.Errorf("cannot ping PostgreSQL: %v", err)
	}

	s := PostgresStorage{
		db: db,
	}
	return &s, nil
}

// ClosePostgres закрывает пул соединений к базе данных PostgreSQL.
func (ps *PostgresStorage) ClosePostgres() {
	ps.db.Close()
}
