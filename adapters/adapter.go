package adapters

type Adapter[T interface{}] interface {
	Read() (*T, error)
	Write(data T) error
}
