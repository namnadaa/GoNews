package memdb

import (
	"GoNews/pkg/storage"
	"sync"
)

// Хранилище данных.
type Store struct {
	mu      sync.RWMutex
	Posts   []storage.Post
	Users   []storage.User
	counter int
}

// Конструктор объекта хранилища.
func NewStore() *Store {
	return &Store{
		Posts:   []storage.Post{},
		Users:   []storage.User{},
		counter: 1,
	}
}
