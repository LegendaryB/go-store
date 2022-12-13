package adapters

type InMemoryAdapter[T interface{}] struct {
	data T
}

func NewInMemoryAdapter[T any]() *InMemoryAdapter[T] {
	return &InMemoryAdapter[T]{}
}

func (adapter *InMemoryAdapter[T]) Read() (*T, error) {
	return &adapter.data, nil
}

func (adapter *InMemoryAdapter[T]) Write(data T) error {
	adapter.data = data

	return nil
}
