package storage

import "time"

// User — структура, представляющая пользователя в in-memory хранилище.
type User struct {
	ID   int
	Name string
}

// Post - публикация.
type Post struct {
	ID          int
	Title       string //
	Content     string //
	AuthorID    int    //
	AuthorName  string
	CreatedAt   time.Time
	PublishedAt time.Time //
}

// Interface задаёт контракт на работу с БД.
type Interface interface {
	GetPosts() ([]Post, error) // получение всех публикаций
	AddPost(Post) error        // создание новой публикации
	UpdatePost(Post) error     // обновление публикации
	DeletePost(Post) error     // удаление публикации по ID
}
