package store

import "github.com/LegendaryB/go-store/adapters"

type Store[T any] struct {
	adapter adapters.Adapter[T]
	Data    T
}

func New[T any](adapter adapters.Adapter[T]) *Store[T] {
	return &Store[T]{
		adapter: adapter,
	}
}

func (store *Store[T]) Read() error {
	data, err := store.adapter.Read()

	if err != nil {
		return err
	}

	store.Data = *data

	return nil
}
