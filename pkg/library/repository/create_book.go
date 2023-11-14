package repository

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"

	"library-api/pkg/library"
)

func (r *repository) CreateBook(book *library.Book) error {
	_, err := r.conn.InsertOne(context.Background(), bson.D{
		{Key: "title", Value: book.Title},
		{Key: "description", Value: book.Description},
		{Key: "writer", Value: book.Writer},
		{Key: "gender", Value: book.Gender},
		{Key: "release_date", Value: book.ReleaseDate},
		{Key: "created_at", Value: book.CreatedAt},
		{Key: "updated_at", Value: book.UpdatedAt},
	})
	if err != nil {
		return err
	}

	return nil
}
