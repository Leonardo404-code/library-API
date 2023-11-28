package repository

import (
	"context"
	"fmt"

	"library-api/pkg/library"
)

func (r *repository) Search(filter *library.Filter) (book []*library.Book, err error) {
	query, err := filter.GenerateQuery()
	if err != nil {
		return nil, err
	}

	collection := r.conn.Database("library").Collection("books")

	result, err := collection.Find(context.Background(), query)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrFindBook, err)
	}

	if err = result.All(context.Background(), &book); err != nil {
		return nil, fmt.Errorf("%w: %v", ErrParseResult, err)
	}

	return book, nil
}
