package main

import (
	"fmt"
	"log"

	"github.com/EronAlves1996/go-key-value-store/storage"
)

func main() {
	kv := *storage.New(storage.JsonStorage, "store.json")

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
