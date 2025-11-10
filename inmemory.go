package main

import "sync"

type InMemoryStorage struct {
	data map[string]string
	mu   sync.RWMutex
}

var _ Storage = &InMemoryStorage{}

func (i *InMemoryStorage) lock() func() {
	i.mu.Lock()
	return func() {
		i.mu.Unlock()
	}
}

func (i *InMemoryStorage) rlock() func() {
	i.mu.RLock()
	return func() {
		i.mu.RUnlock()
	}
}

func (i *InMemoryStorage) internalGet(key string) (string, error) {
	value, ok := i.data[key]
	if !ok {
		return "", errNotFound
	}
	return value, nil
}

func (i *InMemoryStorage) Delete(key string) (string, error) {
	defer i.lock()()

	value, err := i.internalGet(key)
	if err != nil {
		return "", err
	}

	delete(i.data, key)
	return value, nil
}

func (i *InMemoryStorage) Get(key string) (string, error) {
	defer i.rlock()()
	return i.internalGet(key)
}

func (i *InMemoryStorage) Set(key string, value string) error {
	defer i.lock()()
	i.data[key] = value
	return nil
}
