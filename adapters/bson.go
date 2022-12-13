package adapters

import "github.com/LegendaryB/go-store/utils"

type BSONAdapter[T interface{}] struct {
	path string
}

func NewBSONAdapter[T interface{}](path string) *BSONAdapter[T] {
	return &BSONAdapter[T]{
		path: path,
	}
}

func NewBSONAdapterAndCreateFile[T interface{}](path string) (*BSONAdapter[T], error) {
	if err := utils.CreateFileIfNotExist(path); err != nil {
		return nil, err
	}

	return (*BSONAdapter[T])(NewBSONAdapter[T](path)), nil
}
