package storage

import (
	"errors"
	"sync"
)

var ErrNotFound = errors.New("key not found")

type Storage interface {
	Get(key string) (string, error)
	Set(key, value string) error
	Delete(key string) (string, error)
}

type StorageType int

const (
	XmlStorage StorageType = iota
	JsonStorage
	InMemoryStorage
)

func New(t StorageType, ins []Interceptor, f string) *Storage {
	data := make(map[string]string)
	var s Storage
	inMemory := &inMemoryStorage{
		mu:           &sync.RWMutex{},
		data:         data,
		interceptors: ins,
	}

	switch t {
	case XmlStorage:
		s = &xmlStore{
			inMemoryStorage: *inMemory,
			file:            f,
		}
	case JsonStorage:
		s = &jsonStore{
			inMemoryStorage: *inMemory,
			file:            f,
		}
	case InMemoryStorage:
		s = inMemory
	}

	return &s
}
