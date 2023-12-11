package service

import (
	"context"
	"fmt"
	"time"

	"library-api/pkg/library"
)

func (s *service) DeleteBook(requestParams *library.Filter) error {
	if err := requestParams.ValidateParams(); err != nil {
		return fmt.Errorf("%w: %v", ErrInvalidObjectID, err)
	}

	books, err := s.libraryRepo.Search(requestParams)
	if err != nil {
		return fmt.Errorf("%w: %v", ErrGetBook, err)
	}

	if len(books) < 1 {
		return ErrNotFound
	}

	var bookTitle string

	for _, book := range books {
		bookTitle = book.Title
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*50)
	defer cancel()

	if err := s.libraryRepo.DeleteBook(requestParams, s.deleteFromBucket(ctx, bookTitle)); err != nil {
		return fmt.Errorf("%w: %v", ErrDeleteBook, err)
	}

	return nil
}

func (s *service) deleteFromBucket(ctx context.Context, bookTitle string) func() error {
	return func() error {
		return s.storage.DeleteFromBucket(ctx, bookTitle)
	}
}
