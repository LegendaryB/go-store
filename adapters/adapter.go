package adapters

type Adapter[T interface{}] interface {
	Read() T
	Write(data T)
}
