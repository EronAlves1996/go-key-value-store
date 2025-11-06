package main

import (
	"encoding/json"
	"errors"
	"os"
	"sync"
)

type KVStore struct {
	mu   sync.RWMutex
	data map[string]string
	file string
}

func (k *KVStore) save() error {
	f, err := os.OpenFile(k.file, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0o755)
	if err != nil {
		return err
	}
	defer f.Close()

	if err := json.NewEncoder(f).Encode(k.data); err != nil {
		return err
	}

	return nil
}

func (k *KVStore) Set(key, value string) error {
	k.mu.Lock()
	defer k.mu.Unlock()
	defer k.save()
	k.data[key] = value
	return nil
}

func (k *KVStore) Get(key string) (string, error) {
	k.mu.RLock()
	defer k.mu.RUnlock()
	v, ok := k.data[key]

	if !ok {
		return "", errors.New("key not found")
	}

	return v, nil
}
