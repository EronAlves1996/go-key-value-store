package storage

import (
	"encoding/json"
	"fmt"
	"os"
)

type jsonStore struct {
	inMemoryStorage
	file string
}

var _ Storage = &jsonStore{}

func (k *jsonStore) Get(key string) (string, error) {
	return k.inMemoryStorage.Get(key)
}

// Delete implements Storage.
func (k *jsonStore) Delete(key string) (string, error) {
	value, err := k.inMemoryStorage.Get(key)

	if err != nil {
		return "", err
	}

	k.inMemoryStorage.Delete(key)
	if err := k.save(); err != nil {
		k.data[key] = value
		return "", fmt.Errorf("unable to delete key %q: %w", key, err)
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
	// Trying to make keyset atomically
	oldValue, found := k.inMemoryStorage.Get(key)

	var err error
	defer func() {
		if err == nil {
			return
		}

		if found == nil {
			k.inMemoryStorage.Delete(key)
			return
		}

		k.inMemoryStorage.Set(key, oldValue)
	}()

	err = k.inMemoryStorage.Set(key, value)
	if err != nil {
		return fmt.Errorf("unable to save key %q: %w", key, err)
	}

	err = k.save()
	if err != nil {
		return fmt.Errorf("unable to save key %q: %w", key, err)
	}

	return nil
}
