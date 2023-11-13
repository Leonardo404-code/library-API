package library

type Repository interface {
	Search(*Filter) ([]*Book, error)
	CreateBook(*Book) error
}
