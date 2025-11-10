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

func New(t StorageType, f string) *Storage {
	mu := sync.RWMutex{}
	data := make(map[string]string)
	var s Storage

	switch t {
	case XmlStorage:
		s = &xmlStore{
			mu:   &mu,
			data: data,
			file: f,
		}
	case JsonStorage:
		s = &jsonStore{
			mu:   &mu,
			data: data,
			file: f,
		}
	case InMemoryStorage:
		s = &inMemoryStorage{
			data: data,
			mu:   &mu,
		}
	}

	return &s
}
