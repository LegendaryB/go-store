package adapters

type Adapter[T any] interface {
	Read() (*T, error)
	Write(data T) error
}
