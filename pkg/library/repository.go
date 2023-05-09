package library

type Repository interface {
	Search(bookName, bookID string) (*Book, error)
	CreateBook(*Book) error
}
