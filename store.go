package store

import "github.com/LegendaryB/go-store/adapters"

type Store[T any] struct {
	adapter adapters.Adapter[T]
	data    T
}

func New[T any](adapter adapters.Adapter[T]) *Store[T] {
	return &Store[T]{
		adapter: adapter,
	}
}

func NewWithJSONAdapter[T any](opts adapters.JSONAdapterOptions) (*Store[T], error) {
	adapter, err := adapters.NewJSONAdapter[T](opts)

	if err != nil {
		return nil, err
	}

	return &Store[T]{
		adapter: adapter,
	}, nil
}

func NewWithInMemoryAdapter[T any]() *Store[T] {
	adapter := adapters.NewInMemoryAdapter[T]()

	return &Store[T]{
		adapter: adapter,
	}
}

func (store *Store[T]) Read() error {
	data, err := store.adapter.Read()

	if err != nil {
		return err
	}

	store.data = *data

	return nil
}

func (store *Store[T]) Write() error {
	if err := store.adapter.Write(store.data); err != nil {
		return err
	}

	return nil
}
