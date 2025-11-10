package main

import (
	"encoding/xml"
	"errors"
	"os"
	"sync"
)

type XmlStore struct {
	mu   sync.RWMutex
	data map[string]string
	file string
}

var _ Storage = &XmlStore{}

func (x *XmlStore) lock() func() {
	x.mu.Lock()
	return func() {
		x.mu.Unlock()
	}
}

func (x *XmlStore) rlock() func() {
	x.mu.RLock()
	return func() {
		x.mu.RUnlock()
	}
}

func (x *XmlStore) save() error {
	f, err := os.OpenFile(x.file, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0o600)
	if err != nil {
		return err
	}

	defer f.Close()

	if err := xml.NewEncoder(f).Encode(x.data); err != nil {
		return err
	}

	return nil
}

// Delete implements Storage.
func (x *XmlStore) Delete(key string) (string, error) {
	defer x.lock()()
	value, err := x.internalGet(key)

	if err != nil {
		return "", err
	}

	if err := x.save(); err != nil {
		x.data[key] = value
		return "", err
	}

	return value, nil
}

func (x *XmlStore) internalGet(key string) (string, error) {
	value, ok := x.data[key]
	if !ok {
		return "", errNotFound
	}
	return value, nil
}

// Get implements Storage.
func (x *XmlStore) Get(key string) (string, error) {
	defer x.rlock()()
	return x.internalGet(key)
}

// Set implements Storage.
func (x *XmlStore) Set(key string, value string) error {
	defer x.lock()()

	// Trying to make keyset atomically
	oldValue, found := x.data[key]

	var err error
	defer func() {
		if err == nil {
			return
		}

		if !found {
			delete(x.data, key)
			return
		}

		x.data[key] = oldValue
	}()

	x.data[key] = value

	err = x.save()
	if err != nil {
		return errors.New("unable to save key")
	}

	return nil
}
