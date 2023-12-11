package service

import (
	"context"
	"fmt"
	"time"

	"library-api/pkg/library"
)

func (s *service) DownloadBook(filter *library.Filter) error {
	books, err := s.libraryRepo.Search(filter)
	if err != nil {
		return fmt.Errorf("%w: %v", ErrGetBook, err)
	}

	if len(books) == 0 {
		return ErrNotFound
	}

	var bookTitle string

	for _, book := range books {
		bookTitle = book.Title
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*50)
	defer cancel()

	if err = s.storage.DownloadBook(ctx, bookTitle); err != nil {
		return err
	}

	return nil
}
