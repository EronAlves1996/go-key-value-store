package storage

import "sync"

type inMemoryStorage struct {
	data map[string]string
	mu   *sync.RWMutex
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

func (i *inMemoryStorage) internalGet(key string) (string, error) {
	value, ok := i.data[key]
	if !ok {
		return "", ErrNotFound
	}
	return value, nil
}

func (i *inMemoryStorage) Delete(key string) (string, error) {
	defer i.lock()()

	value, err := i.internalGet(key)
	if err != nil {
		return "", err
	}

	delete(i.data, key)
	return value, nil
}

func (i *inMemoryStorage) Get(key string) (string, error) {
	defer i.rlock()()
	return i.internalGet(key)
}

func (i *inMemoryStorage) Set(key string, value string) error {
	defer i.lock()()
	i.data[key] = value
	return nil
}
