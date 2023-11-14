package service

import (
	"fmt"
	"time"

	"library-api/pkg/library"
)

func (s *service) SendBook(book *library.Book) (*library.Book, error) {
	if len(book.Title) < 3 || len(book.Title) > 80 {
		return nil, fmt.Errorf(
			"%w: the field 'name' must be between 3 and 80 characters long",
			ErrInvalidName,
		)
	}

	if len(book.Description) < 8 {
		return nil, fmt.Errorf(
			"%w: field 'description' must be longer than 8 characters",
			ErrInvalidDescription,
		)
	}

	if _, err := time.Parse(onlyDateLayout, book.ReleaseDate); err != nil {
		return nil, fmt.Errorf("%w: %v", ErrInvalidDate, err)
	}

	bookModel := &library.Book{
		Title:       book.Title,
		Description: book.Description,
		Writer:      book.Writer,
		Gender:      book.Gender,
		ReleaseDate: book.ReleaseDate,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	if err := s.libraryRepo.CreateBook(bookModel); err != nil {
		return nil, fmt.Errorf("%w: %v", ErrCreateBook, err)
	}

	return bookModel, nil
}
