package main

import (
	"errors"
	"fmt"
	"log"
	"sync"
)

type Storage interface {
	Get(key string) (string, error)
	Set(key, value string) error
	Delete(key string) (string, error)
}

var errNotFound = errors.New("key not found")

func main() {
	kv := JsonStore{
		mu:   sync.RWMutex{},
		data: make(map[string]string),
		file: "store.json",
	}

	if err := kv.Set("test", "teakslfjaskfas"); err != nil {
		log.Fatal(err)
	}

	if err := kv.Set("aklsflafas", "fkasfhaklfkjashfslf"); err != nil {
		log.Fatal(err)
	}

	value, err := kv.Get("test")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(value)
}
