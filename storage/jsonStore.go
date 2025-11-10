package storage

import (
	"encoding/json"
	"errors"
	"os"
	"sync"
)

type jsonStore struct {
	mu   *sync.RWMutex
	data map[string]string
	file string
}

var _ Storage = &jsonStore{}

func (k *jsonStore) internalGet(key string) (string, error) {
	v, ok := k.data[key]

	if !ok {
		return "", ErrNotFound
	}

	return v, nil
}

func (k *jsonStore) Get(key string) (string, error) {
	k.mu.RLock()
	defer k.mu.RUnlock()
	return k.internalGet(key)
}

// Delete implements Storage.
func (k *jsonStore) Delete(key string) (string, error) {
	k.mu.Lock()
	defer k.mu.Unlock()
	value, err := k.internalGet(key)

	if err != nil {
		return "", err
	}

	delete(k.data, key)
	if err := k.save(); err != nil {
		k.data[key] = value
		return "", errors.New("unable to delete key")
	}

	return value, nil
}

func (k *jsonStore) save() error {
	f, err := os.OpenFile(k.file, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0o600)
	if err != nil {
		return err
	}
	defer f.Close()

	if err := json.NewEncoder(f).Encode(k.data); err != nil {
		return err
	}

	return nil
}

func (k *jsonStore) Set(key, value string) error {
	k.mu.Lock()
	defer k.mu.Unlock()

	// Trying to make keyset atomically
	oldValue, found := k.data[key]

	var err error
	defer func() {
		if err == nil {
			return
		}

		if !found {
			delete(k.data, key)
			return
		}

		k.data[key] = oldValue
	}()

	k.data[key] = value

	err = k.save()
	if err != nil {
		return errors.New("unable to save key")
	}

	return nil
}
