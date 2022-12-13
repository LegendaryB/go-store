package main

import (
	"fmt"

	"github.com/LegendaryB/go-store"
	"github.com/LegendaryB/go-store/adapters"
)

type User struct {
	Name string
}

func main() {
	opts := adapters.JSONAdapterOptions{
		Path:                     "db.json",
		CreateFileWhenNotPresent: true,
		UsePrettyPrint:           true,
	}

	adapter, err := adapters.NewJSONAdapter[User](opts)

	if err != nil {
		panic(err)
	}

	store := store.New[User](adapter)

	if err := store.Read(); err != nil {
		panic(err)
	}

	fmt.Printf(store.Data.Name)
}
