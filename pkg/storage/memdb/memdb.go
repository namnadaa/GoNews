package memdb

import (
	"GoNews/pkg/storage"
	"sync"
)

// Хранилище данных.
type MemoryStorage struct {
	mu      sync.RWMutex
	Posts   []storage.Post
	Users   []storage.User
	counter int
}

// Конструктор объекта хранилища.
func NewMemoryStorage() *MemoryStorage {
	return &MemoryStorage{
		Posts:   []storage.Post{},
		Users:   []storage.User{},
		counter: 1,
	}
}
