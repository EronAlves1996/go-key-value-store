package storage

import (
	"sync"

	"github.com/EronAlves1996/go-key-value-store/errors"
)

type inMemoryStorage struct {
	data         map[string]string
	interceptors []Interceptor
	mu           *sync.RWMutex
}

var _ Storage = &inMemoryStorage{}

func (i *inMemoryStorage) lock() func() {
	i.mu.Lock()
	return func() {
		i.mu.Unlock()
	}
}

func (i *inMemoryStorage) rlock() func() {
	i.mu.RLock()
	return func() {
		i.mu.RUnlock()
	}
}

func (i *inMemoryStorage) runInterceptors(method string, inc *InterceptorContext) {
	for _, in := range i.interceptors {
		in.Intercept(method, inc)
	}
}

func (i *inMemoryStorage) interceptorLessGet(inc *InterceptorContext) {
	value, ok := i.data[inc.Key]

	if !ok {
		inc.Err = ErrNotFound
		return
	}

	inc.Value = value
}

func (i *inMemoryStorage) internalGet(inc *InterceptorContext) {
	defer i.runInterceptors("GET", inc)
	i.interceptorLessGet(inc)
}

func (i *inMemoryStorage) internalSet(inc *InterceptorContext) {
	defer i.runInterceptors("SET", inc)
	if inc.Key == "" {
		inc.Err = &errors.InvalidKeyError{
			Key: inc.Key,
		}
		return
	}
	i.data[inc.Key] = inc.Value
}

func (i *inMemoryStorage) internalDelete(inc *InterceptorContext) {
	defer i.runInterceptors("DELETE", inc)

	i.interceptorLessGet(inc)

	if inc.Err != nil {
		return
	}

	delete(i.data, inc.Key)
}

func (i *inMemoryStorage) Delete(key string) (string, error) {
	defer i.lock()()
	inc := InterceptorContext{
		Key: key,
	}

	i.internalDelete(&inc)

	return inc.Value, inc.Err
}

func (i *inMemoryStorage) Get(key string) (string, error) {
	defer i.rlock()()

	inc := InterceptorContext{
		Key: key,
	}
	i.internalGet(&inc)

	return inc.Value, inc.Err
}

func (i *inMemoryStorage) Set(key string, value string) error {
	defer i.lock()()

	inc := InterceptorContext{
		Key:   key,
		Value: value,
	}

	i.internalSet(&inc)

	return inc.Err
}
