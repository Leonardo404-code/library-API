package library

type Repository interface {
	Search(bookName, bookID string) (*Book, error)
}
