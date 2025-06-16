package mongo

import (
	"GoNews/pkg/storage"
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
)

// GetPosts возвращает список всех постов из MongoDB.
func (ms *MongoStorage) GetPosts() ([]storage.Post, error) {
	collection := ms.db.Database(ms.databaseName).Collection(ms.collectionName)
	filter := bson.D{}
	cur, err := collection.Find(context.Background(), filter)
	if err != nil {
		return nil, fmt.Errorf("failed to find post: %v", err)
	}
	defer cur.Close(context.Background())

	var data []storage.Post
	for cur.Next(context.Background()) {
		var post storage.Post
		err := cur.Decode(&post)
		if err != nil {
			return nil, fmt.Errorf("failed to decode post: %v", err)
		}
		data = append(data, post)
	}
	return data, cur.Err()
}

// AddPost добавляет новый пост в базу данных MongoDB.
func (ms *MongoStorage) AddPost(sp storage.Post) error {
	collection := ms.db.Database(ms.databaseName).Collection(ms.collectionName)
	_, err := collection.InsertOne(context.Background(), sp)
	if err != nil {
		return fmt.Errorf("failed to insert post: %v", err)
	}
	return nil
}

// UpdatePost обновляет пост в базе данных MongoDB.
func (ms *MongoStorage) UpdatePost(sp storage.Post) error {
	collection := ms.db.Database(ms.databaseName).Collection(ms.collectionName)

	filter := bson.M{"id": sp.ID}
	update := bson.M{
		"$set": bson.M{
			"title":       sp.Title,
			"content":     sp.Content,
			"authorid":    sp.AuthorID,
			"authorname":  sp.AuthorName,
			"publishedat": sp.PublishedAt,
		},
	}

	res, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return fmt.Errorf("failed to update post: %v", err)
	}

	if res.MatchedCount == 0 {
		return fmt.Errorf("post with ID %d not found", sp.ID)
	}
	return nil
}

// DeletePost - удаляет пост в базе данных MongoDB.
func (ms *MongoStorage) DeletePost(sp storage.Post) error {
	collection := ms.db.Database(ms.databaseName).Collection(ms.collectionName)

	filter := bson.M{"id": sp.ID}

	res, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		return fmt.Errorf("failed to delete post: %v", err)
	}

	if res.DeletedCount == 0 {
		return fmt.Errorf("post with ID %d not found", sp.ID)
	}
	return nil
}
