package main

import (
	"encoding/json"
	"errors"
	"os"
	"sync"
)

type JsonStore struct {
	mu   sync.RWMutex
	data map[string]string
	file string
}

var _ Storage = &JsonStore{}

func (k *JsonStore) internalGet(key string) (string, error) {
	v, ok := k.data[key]

	if !ok {
		return "", errNotFound
	}

	return v, nil
}

func (k *JsonStore) Get(key string) (string, error) {
	k.mu.RLock()
	defer k.mu.RUnlock()
	return k.internalGet(key)
}

// Delete implements Storage.
func (k *JsonStore) Delete(key string) (string, error) {
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

func (k *JsonStore) save() error {
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

func (k *JsonStore) Set(key, value string) error {
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
