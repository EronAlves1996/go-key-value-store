package main

import (
	"errors"
	"fmt"
	"log"

	internal_errors "github.com/EronAlves1996/go-key-value-store/errors"
	"github.com/EronAlves1996/go-key-value-store/storage"
)

type LoggingInterceptor struct{}

// Intercept implements storage.Interceptor.
func (l LoggingInterceptor) Intercept(method string, i *storage.InterceptorContext) {
	fmt.Printf("%s\t%s\t%s\t%s\n", method, i.Key, i.Value, i.Err)
}

var _ storage.Interceptor = LoggingInterceptor{}

func main() {
	kv := storage.New(storage.XmlStorage, []storage.Interceptor{
		LoggingInterceptor{},
	}, "store.xml")

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

	if err := kv.Set("", "teste"); err != nil {
		if errors.Is(err, &internal_errors.InvalidKeyError{}) {
			fmt.Println("error is invalid key error")
		}

		if errors.As(err, &internal_errors.InvalidKeyError{}) {
			cast, _ := err.(internal_errors.InvalidKeyError)
			fmt.Printf("invalid key error: %v\n", cast.Key)
		}
	}

	fmt.Println(value)
}
