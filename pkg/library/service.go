package library

type Services interface {
	SendBook(*Book) (*Book, error)
}
