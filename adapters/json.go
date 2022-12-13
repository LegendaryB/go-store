package adapters

import "github.com/LegendaryB/go-store/utils"

type JSONAdapter[T interface{}] struct {
	path string
}

func NewJSONAdapter[T interface{}](path string) *JSONAdapter[T] {
	return &JSONAdapter[T]{
		path: path,
	}
}

func NewJSONAdapterAndCreateFile[T interface{}](path string) (*JSONAdapter[T], error) {
	if err := utils.CreateFileIfNotExist(path); err != nil {
		return nil, err
	}

	return (*JSONAdapter[T])(NewJSONAdapter[T](path)), nil
}
