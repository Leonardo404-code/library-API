package library

type Services interface {
	GetBooks(bookName, bookUUID string) (*Book, error)
}
