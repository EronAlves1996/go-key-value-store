package main

import "sync"

type KVStore struct {
	mu   sync.Mutex
	data map[string]string
	file string
}
