package repository

import (
	"context"

	"library-api/pkg/library"
)

func (r *repository) DeleteBook(filter *library.Filter) error {
	query := filter.GenerateQuery()
	collection := r.conn.Database(libraryDB).Collection(booksColl)

	if _, err := collection.DeleteOne(context.TODO(), query); err != nil {
		return err
	}

	return nil
}
