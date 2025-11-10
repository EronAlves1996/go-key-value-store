package main

import (
	"fmt"
	"log"

	"github.com/EronAlves1996/go-key-value-store/storage"
)

type LoggingInterceptor struct{}

// Intercept implements storage.Interceptor.
func (l LoggingInterceptor) Intercept(method string, i *storage.InterceptorContext) {
	fmt.Printf("%s\t%s\t%s\t%s\n", method, i.Key, i.Value, i.Err)
}

var _ storage.Interceptor = LoggingInterceptor{}

func main() {
	kv := *storage.New(storage.XmlStorage, []storage.Interceptor{
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

	fmt.Println(value)
}
