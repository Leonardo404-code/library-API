package repository

import (
	"context"
	"fmt"
	"library-api/pkg/library"
	"mime/multipart"

	"go.mongodb.org/mongo-driver/bson"
)

func (r *repository) CreateBook(
	book *library.Book,
	bookFile multipart.File,
	upload func() error,
) error {
	session, err := r.conn.StartSession()
	if err != nil {
		return fmt.Errorf("%w: %v", ErrStartSession, err)
	}

	defer session.EndSession(context.TODO())

	if err = session.StartTransaction(); err != nil {
		return fmt.Errorf("%w: %v", ErrStartSession, err)
	}

	collection := r.conn.Database("library").Collection("books")

	if _, err := collection.InsertOne(context.Background(), bson.D{
		{Key: "title", Value: book.Title},
		{Key: "description", Value: book.Description},
		{Key: "writer", Value: book.Writer},
		{Key: "gender", Value: book.Gender},
		{Key: "release_date", Value: book.ReleaseDate},
		{Key: "book_url", Value: book.BookURL},
		{Key: "created_at", Value: book.CreatedAt},
		{Key: "updated_at", Value: book.UpdatedAt},
	}); err != nil {
		return err
	}

	if err = upload(); err != nil {
		return err
	}

	if err = session.CommitTransaction(context.TODO()); err != nil {
		return fmt.Errorf("%w: %v", ErrCommitTransaction, err)
	}

	return nil
}
