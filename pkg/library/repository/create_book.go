package repository

import (
	"context"

	"library-api/pkg/library"

	"go.mongodb.org/mongo-driver/bson"
)

func (r *repository) CreateBook(book *library.Book) error {
	_, err := r.conn.InsertOne(context.Background(), bson.D{
		{Key: "name", Value: book.Name},
		{Key: "description", Value: book.Description},
		{Key: "writer", Value: book.Writer},
		{Key: "release_date", Value: book.ReleaseDate},
		{Key: "created_at", Value: book.CreatedAt},
		{Key: "updated_at", Value: book.UpdatedAt},
	})
	if err != nil {
		return err
	}

	return nil
}
