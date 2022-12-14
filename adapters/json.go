package adapters

import (
	"encoding/json"
	"os"
	"path/filepath"
)

type JSONAdapterOptions struct {
	Path                     string
	CreateFileWhenNotPresent bool
	UsePrettyPrint           bool
}

type JSONAdapter[T any] struct {
	path                     string
	createFileWhenNotPresent bool
	usePrettyPrint           bool
}

func NewJSONAdapter[T any](opts JSONAdapterOptions) (*JSONAdapter[T], error) {
	adapter := &JSONAdapter[T]{
		path:                     opts.Path,
		createFileWhenNotPresent: opts.CreateFileWhenNotPresent,
		usePrettyPrint:           opts.UsePrettyPrint,
	}

	if adapter.createFileWhenNotPresent {
		if err := adapter.createEmptyJSONFile(); err != nil {
			return nil, err
		}
	}

	return adapter, nil
}

func (adapter *JSONAdapter[T]) createEmptyJSONFile() error {
	directoryPath := filepath.Dir(adapter.path)

	if err := os.MkdirAll(directoryPath, os.ModePerm); err != nil {
		return err
	}

	if _, err := os.Stat(adapter.path); os.IsNotExist(err) {
		os.WriteFile(adapter.path, []byte(`{}`), os.ModePerm)
	}

	return nil
}

func (adapter *JSONAdapter[T]) Read() (*T, error) {
	content, err := os.ReadFile(adapter.path)

	if err != nil {
		return nil, err
	}

	var data T

	err = json.Unmarshal(content, &data)

	if err != nil {
		return nil, err
	}

	return &data, nil
}

func (adapter *JSONAdapter[T]) Write(data T) error {
	var buffer []byte
	var err error

	if adapter.usePrettyPrint {
		buffer, err = json.MarshalIndent(data, "", "\t")
	} else {
		buffer, err = json.Marshal(data)
	}

	if err != nil {
		return err
	}

	if err = os.WriteFile(adapter.path, buffer, os.ModePerm); err != nil {
		return err
	}

	return nil
}
