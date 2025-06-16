package memdb

import (
	"GoNews/pkg/storage"
	"fmt"
	"time"
)

// Posts - возвращает список всех постов.
func (s *Store) GetPosts() ([]storage.Post, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	posts := make([]storage.Post, len(s.Posts))

	for i, p := range s.Posts {
		posts[i] = p
		for _, u := range s.Users {
			if u.ID == p.AuthorID {
				posts[i].AuthorName = u.Name
				break
			}
		}
	}

	return posts, nil
}

// AddPost добавляет новый пост в память.
func (s *Store) AddPost(sp storage.Post) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	var authorName string
	for _, u := range s.Users {
		if u.ID == sp.AuthorID {
			authorName = u.Name
			break
		}
	}

	if authorName == "" {
		return fmt.Errorf("author with ID %d not found", sp.AuthorID)
	}

	publishedAt := sp.PublishedAt
	if publishedAt.IsZero() {
		publishedAt = time.Now()
	}

	newPost := storage.Post{
		ID:          s.counter,
		Title:       sp.Title,
		Content:     sp.Content,
		AuthorID:    sp.AuthorID,
		AuthorName:  authorName,
		CreatedAt:   time.Now(),
		PublishedAt: publishedAt,
	}

	s.Posts = append(s.Posts, newPost)
	s.counter++

	return nil
}

// UpdatePost обновляет пост.
func (s *Store) UpdatePost(sp storage.Post) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	var authorName string
	for _, u := range s.Users {
		if u.ID == sp.AuthorID {
			authorName = u.Name
			break
		}
	}

	if authorName == "" {
		return fmt.Errorf("author with ID %d not found", sp.AuthorID)
	}

	publishedAt := sp.PublishedAt
	if publishedAt.IsZero() {
		publishedAt = time.Now()
	}

	for i, p := range s.Posts {
		if p.ID == sp.ID {
			s.Posts[i].Title = sp.Title
			s.Posts[i].Content = sp.Content
			s.Posts[i].AuthorID = sp.AuthorID
			s.Posts[i].AuthorName = authorName
			s.Posts[i].PublishedAt = publishedAt
			return nil
		}
	}

	return fmt.Errorf("post with ID %d not found", sp.ID)
}

// DeletePost - удаляет пост.
func (s *Store) DeletePost(sp storage.Post) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	for i, p := range s.Posts {
		if p.ID == sp.ID {
			s.Posts = append(s.Posts[:i], s.Posts[i+1:]...)
			return nil
		}
	}

	return fmt.Errorf("post with ID %d not found", sp.ID)
}
