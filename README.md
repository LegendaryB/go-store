<div align="center">

<h1>go-store</h1>

[![forthebadge](https://forthebadge.com/images/badges/fuck-it-ship-it.svg)](https://forthebadge.com)
[![forthebadge](https://forthebadge.com/images/badges/made-with-go.svg)](https://forthebadge.com)

[![GitHub license](https://img.shields.io/github/license/LegendaryB/go-store.svg?longCache=true&style=flat-square)](https://github.com/LegendaryB/go-store/blob/main/LICENSE.txt)

<sub>Built with ❤︎ by LegendaryB</sub>

</div><br>

## 🎯 Features

- Lightweight
- Pure Golang, no external dependencies
- Minimalistic
- Extendable

## 🚀 Install

```
go get github.com/LegendaryB/go-store
```

## 📝 Usage

```go
package main

import (
	"github.com/LegendaryB/go-store"
	"github.com/LegendaryB/go-store/adapters"
)

type Blog struct {
	Articles []Article
}

type Article struct {
	Title   string
	Content string
	Author  Author
}

type Author struct {
	Name string
}

func main() {
        // Configure the options for the JSON Adapter
	opts := adapters.JSONAdapterOptions{
		Path:                     "db.json",
		CreateFileWhenNotPresent: true,
		UsePrettyPrint:           true,
	}

        // Create a new instance of the JSON Adapter type
	adapter, err := adapters.NewJSONAdapter[Blog](opts)

	if err != nil {
		panic(err)
	}

        // Create a new instance of the Store type and configure it to use the JSON Adapter type instance
        // Equal to using: store.NewWithJSONAdapter(opts)
	store := store.New[Blog](adapter)

        // Read the data from disk
	if err := store.Read(); err != nil {
		panic(err)
	}

        // Write the data to disk
	defer store.Write()

        // Use the native Go api to append a new article into the articles array
	store.Data.Articles = append(store.Data.Articles, Article{
		Title:   "Lorem Ipsum",
		Content: "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Aliquam mauris tellus, interdum in neque in, aliquam pharetra tellus. Vivamus mollis facilisis lacinia. In maximus auctor volutpat. Phasellus vel elit justo. Sed mattis elit vitae purus commodo vehicula. Curabitur tristique lacus sed blandit suscipit. Aenean lobortis vitae ligula eget rutrum. Phasellus ut iaculis justo. Praesent molestie aliquam justo, ac pellentesque nisl luctus ut. Maecenas dictum aliquet justo, sollicitudin molestie nisl accumsan blandit. Nunc placerat erat id dui ultricies mollis. Morbi venenatis facilisis sodales. Donec eget risus urna. Maecenas pulvinar felis urna, vitae molestie metus dictum sed. Aenean nec vulputate erat.",
		Author: Author{
			Name: "N/A",
		},
	})
}
```

## 🔌 Adapters

### JSON adapter

This adapter is used for reading and writing data to a JSON file. It can be used like this:

```go
func main() {
        // Configure the options for the JSON Adapter
	opts := adapters.JSONAdapterOptions{
		Path:                     "db.json",
		CreateFileWhenNotPresent: true,
		UsePrettyPrint:           true,
	}

        // Create a new instance of the JSON Adapter type
	adapter, err := adapters.NewJSONAdapter[Blog](opts)

	if err != nil {
		panic(err)
	}

        // Use it
	store := store.New[Blog](adapter)
}
```

or

```go
func main() {
        // Configure the options for the JSON Adapter
	opts := adapters.JSONAdapterOptions{
		Path:                     "db.json",
		CreateFileWhenNotPresent: true,
		UsePrettyPrint:           true,
	}

        // Use it
	store, err := store.NewWithJSONAdapter[Blog](opts)

        if err != nil {
		panic(err)
	}
}
```

### In memory adapter

This adapter can be used for quick testing or unit tests. After the process exits the data is lost. It can be used like this:

```go
func main() {
        // Create a new instance of the InMemoryAdapter type
	adapter, err := adapters.NewInMemoryAdapter[Blog]()

	if err != nil {
		panic(err)
	}

        // Use it
	store := store.New[Blog](adapter)
}
```

or

```go
func main() {
        // Use it
	store, err := store.NewWithInMemoryAdapter[Blog](opts)

        if err != nil {
		panic(err)
	}
}
```

### Writing a custom adapter

To implement a custom provider you only need to implement the `Adapter` interface:

```go
type Adapter[T any] interface {
	Read() (*T, error)
	Write(data T) error
}
```

Afterwards you can use your custom provider like this:

```go
func main() {
        // Create a new instance of your custom provider type
	adapter, err := adapters.NewMyCustomAdapter[Blog]()

	if err != nil {
		panic(err)
	}

        // Use it
	store := store.New[Blog](adapter)
}
```
