﻿<h1 align="center">gostore</h1><div align="center">

[![forthebadge](https://forthebadge.com/images/badges/fuck-it-ship-it.svg)](https://forthebadge.com)
[![forthebadge](https://forthebadge.com/images/badges/made-with-go.svg)](https://forthebadge.com)

[![GitHub license](https://img.shields.io/github/license/LegendaryB/gogdl-ng.svg?longCache=true&style=flat-square)](https://github.com/LegendaryB/gogdl-ng/blob/main/LICENSE)

<sub>Built with ❤︎ by LegendaryB</sub>

gostore is a small and lightweight embedded database.  
</div><br>

## Getting started

### Creating a store instance
To create a store instance use the `gostore.New("gostore.json")` like so:
```go
store, err := gostore.New("gostore.json", false)

if err != nil {
    fmt.Print(err)
}
```

### Creating a collection
To create a collection use the `store.CreateCollection("awesomeCollection")` method like so:
```go
store, err := gostore.New("gostore.json")

if err != nil {
    fmt.Print(err)
}

// returns a new or existing collection
collection, err := store.CreateCollection("awesomeCollection")
```

### Getting a collection
To get a existing collection use the `store.GetCollection("awesomeCollection")` method like so:
```go
store, err := gostore.New("gostore.json")

if err != nil {
    fmt.Print(err)
}

collection, err := store.GetCollection("awesomeCollection")
```

### Delete a collection
To delete a collection use the `store.DeleteCollection("awesomeCollection")` method like so:
```go
store, err := gostore.New("gostore.json")

if err != nil {
    fmt.Print(err)
}

// returns a new or existing collection
collection, err := store.DeleteCollection("awesomeCollection")
```

### Getting the collection names
To get all collection names use the `store.GetCollectionNames()` method like so:
```go
store, err := gostore.New("gostore.json")

if err != nil {
    fmt.Print(err)
}

cNames, err := store.GetCollectionNames()
```

### Add a item to the collection

### Get a item from the collection
To get a item from a collection use the `collection.Get(index, any)` method like so:
```go
awesomeStruct := &AwesomeStruct{}

err = collection.Get(0, awesomeStruct)

if err != nil {
    fmt.Print(awesomeStruct.Property)
}
```


### Set (update) a item in the collection 

### Delete a item from the collection

### Finding a item in the collection

### Get all items from the collection
