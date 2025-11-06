package main

import "sync"

type KVStore struct {
	mu   sync.Mutex
	data map[string]string
	file string
}

func (k *KVStore) Set(key, value string) error {
	k.mu.Lock()
	defer k.mu.Unlock()
	k.data[key] = value
	return nil
}
