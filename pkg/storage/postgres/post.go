package postgres

import (
	"GoNews/pkg/storage"
	"context"
	"fmt"
)

// GetPosts возвращает список всех постов из PostgreSQL.
func (ps *PostgresStorage) GetPosts() ([]storage.Post, error) {
	rows, err := ps.db.Query(context.Background(), `
	SELECT
		posts.id,
		posts.title,
		posts.content,
		posts.author_id,
		authors.name,
		posts.created_at,
		posts.published_at
	FROM
		posts
	JOIN authors ON posts.author_id = authors.id
	ORDER BY created_at;
`,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query for GetPosts: %v", err)
	}
	defer rows.Close()

	var posts []storage.Post
	for rows.Next() {
		var p storage.Post
		err = rows.Scan(
			&p.ID,
			&p.Title,
			&p.Content,
			&p.AuthorID,
			&p.AuthorName,
			&p.CreatedAt,
			&p.PublishedAt,
		)

		if err != nil {
			return nil, fmt.Errorf("failed to scan post row: %v", err)
		}

		posts = append(posts, p)
	}
	return posts, rows.Err()
}

// AddPost добавляет новый пост в базу данных PostgreSQL.
func (ps *PostgresStorage) AddPost(sp storage.Post) error {
	_, err := ps.db.Exec(context.Background(), `
	INSERT INTO posts (title, content, author_id, published_at)
	VALUES ($1, $2, $3, $4);
	`,
		sp.Title, sp.Content, sp.AuthorID, sp.PublishedAt,
	)

	if err != nil {
		return fmt.Errorf("failed to create post: %v", err)
	}
	return nil
}

// UpdatePost обновляет пост в базе данных PostgreSQL.
func (ps *PostgresStorage) UpdatePost(sp storage.Post) error {
	res, err := ps.db.Exec(context.Background(), `
	UPDATE posts
	SET title = $1, content = $2, author_id = $3, published_at = $4
	WHERE id = $5;
	`,
		sp.Title, sp.Content, sp.AuthorID, sp.ID,
	)

	if err != nil {
		return fmt.Errorf("failed to update post: %v", err)
	}

	if res.RowsAffected() == 0 {
		return fmt.Errorf("post with ID %d not found", sp.ID)
	}
	return nil
}

// DeletePost - удаляет пост в базе данных PostgreSQL.
func (ps *PostgresStorage) DeletePost(sp storage.Post) error {
	res, err := ps.db.Exec(context.Background(), `
	DELETE FROM posts 
	WHERE id = $1;
	`,
		sp.ID,
	)

	if err != nil {
		return fmt.Errorf("failed to delete post: %v", err)
	}

	if res.RowsAffected() == 0 {
		return fmt.Errorf("post with ID %d not found", sp.ID)
	}
	return nil
}
