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
	opts := adapters.JSONAdapterOptions{
		Path:                     "db.json",
		CreateFileWhenNotPresent: true,
		UsePrettyPrint:           true,
	}

	adapter, err := adapters.NewJSONAdapter[Blog](opts)

	if err != nil {
		panic(err)
	}

	store := store.New[Blog](adapter)

	if err := store.Read(); err != nil {
		panic(err)
	}

	defer store.Write()

	store.Data.Articles = append(store.Data.Articles, Article{
		Title:   "Lorem Ipsum",
		Content: "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Aliquam mauris tellus, interdum in neque in, aliquam pharetra tellus. Vivamus mollis facilisis lacinia. In maximus auctor volutpat. Phasellus vel elit justo. Sed mattis elit vitae purus commodo vehicula. Curabitur tristique lacus sed blandit suscipit. Aenean lobortis vitae ligula eget rutrum. Phasellus ut iaculis justo. Praesent molestie aliquam justo, ac pellentesque nisl luctus ut. Maecenas dictum aliquet justo, sollicitudin molestie nisl accumsan blandit. Nunc placerat erat id dui ultricies mollis. Morbi venenatis facilisis sodales. Donec eget risus urna. Maecenas pulvinar felis urna, vitae molestie metus dictum sed. Aenean nec vulputate erat.",
		Author: Author{
			Name: "N/A",
		},
	})
}
